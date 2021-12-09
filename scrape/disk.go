package scrape

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	prommodel "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"

	"github.com/cqroot/openstack-swift-dashboard/models"
)

type hdPair struct {
	host   string
	device string
}

func ScrapeDisk(target models.Target) {
	log.Info().Str("target", target.Name).Msg("Start scraping disk")

	diskMap := make(map[hdPair]*models.Disk)

	totalDisk := models.TotalDisk{
		Target: target.ID,
		Date:   cast.ToInt(time.Now().Local().Format("20060102")),
	}
	disks := make([]models.Disk, 0)

	mf, err := parseMF("http://" + target.Endpoint + "/metrics?collect=disk")
	if err != nil {
		panic(err)
	}
	for k, v := range mf {
		if strings.HasPrefix(k, "swift_disk") {
			if strings.HasPrefix(k, "swift_disk_total") {
				parseTotalDisk(k, v, &totalDisk)

			} else {
				parseDisk(k, v, diskMap)
			}
		}
	}
	for _, d := range diskMap {
		d.Target = target.ID
		disks = append(disks, *d)
	}

	models.UpdateDisks(&disks)
	models.UpdateTotalDisk(&totalDisk)

	log.Info().Str("target", target.Name).Msg("Finish scraping disk")
}

func parseDisk(metricName string, metric *prommodel.MetricFamily, diskMap map[hdPair]*models.Disk) {
	if *metric.Type != prommodel.MetricType_GAUGE {
		return
	}
	for _, metric := range metric.Metric {
		key := hdPair{}
		for _, label := range metric.Label {
			switch *label.Name {
			case "host":
				key.host = *label.Value
			case "device":
				key.device = *label.Value
			}
		}
		if diskMap[key] == nil {
			diskMap[key] = &models.Disk{Host: key.host, Device: key.device}
		}
		switch metricName {
		case "swift_disk_avail_bytes":
			diskMap[key].Avail = int64(*metric.Gauge.Value)
		case "swift_disk_used_bytes":
			diskMap[key].Used = int64(*metric.Gauge.Value)
		case "swift_disk_size_bytes":
			diskMap[key].Size = int64(*metric.Gauge.Value)
		case "swift_disk_usage_bytes":
			diskMap[key].Usage = *metric.Gauge.Value
		}
	}
}

func parseTotalDisk(metricName string, metric *prommodel.MetricFamily, totalDisk *models.TotalDisk) {
	if *metric.Type != prommodel.MetricType_GAUGE {
		return
	}
	for _, metric := range metric.Metric {
		switch metricName {
		case "swift_disk_total_avail_bytes":
			totalDisk.TotalAvail = int64(*metric.Gauge.Value)
		case "swift_disk_total_used_bytes":
			totalDisk.TotalUsed = int64(*metric.Gauge.Value)
		case "swift_disk_total_size_bytes":
			totalDisk.TotalSize = int64(*metric.Gauge.Value)
		}
	}
}

func parseMF(url string) (map[string]*prommodel.MetricFamily, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ParseMF - Get: %w", err)
	}
	defer resp.Body.Close()

	var parser expfmt.TextParser
	mf, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ParseMF - TextToMetricFamilies: %w", err)
	}
	return mf, nil
}
