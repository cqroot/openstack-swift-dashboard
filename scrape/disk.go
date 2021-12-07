package scrape

import (
	"fmt"
	"net/http"
	"strings"

	prommodel "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/rs/zerolog/log"

	"github.com/cqroot/openstack-swift-dashboard/models"
)

type hdPair struct {
	host   string
	device string
}

func ScrapeDisk() error {
	log.Info().Str("Target", "disk").Msg("Start scraping")

	targets, err := models.TargetList()
	if err != nil {
		return err
	}
	for _, target := range targets {
		updateTargetDisks(target)
	}

	log.Warn().Str("Target", "disk").Msg("Finish scraping")
	return nil
}

func updateTargetDisks(target models.Target) {
	diskMap := make(map[hdPair]*models.Disk)
	disks := make([]models.Disk, 0)

	mf, err := parseMF("http://" + target.Endpoint + "/metrics?collect=disk")
	if err != nil {
		panic(err)
	}
	for k, v := range mf {
		if strings.HasPrefix(k, "swift_disk") && !strings.HasPrefix(k, "swift_disk_total") {
			if *v.Type != prommodel.MetricType_GAUGE {
				continue
			}
			for _, metric := range v.Metric {
				key := hdPair{}
				for _, label := range metric.Label {
					switch *label.Name {
					case "host":
						key.host = *label.Value
					case "device":
						key.device = *label.Value
					}
				}
				if key.host == "" {
					fmt.Printf("%+v\n", k)
				}
				if diskMap[key] == nil {
					diskMap[key] = &models.Disk{Host: key.host, Device: key.device}
				}
				switch k {
				case "swift_disk_avail_bytes":
					diskMap[key].Avail = int64(*metric.Gauge.Value)
				case "swift_disk_used_bytes":
					diskMap[key].Used = int64(*metric.Gauge.Value)
				case "swift_disk_size_bytes":
					diskMap[key].Size = int64(*metric.Gauge.Value)
				case "swift_disk_usage_bytes":
					diskMap[key].Usage = int64(*metric.Gauge.Value)
				}
			}
		}
	}
	for _, d := range diskMap {
		disks = append(disks, *d)
	}
	models.UpdateDisks(&disks)
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
