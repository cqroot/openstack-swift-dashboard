package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cqroot/openstack-swift-dashboard/databases"
	"github.com/cqroot/openstack-swift-dashboard/internal"
	"github.com/cqroot/openstack-swift-dashboard/models"
	"github.com/cqroot/openstack-swift-dashboard/scrape"
)

var (
	verbose bool
	dsn     string
)

var rootCmd = &cobra.Command{
	Use:   "scrape",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitZerolog()
		databases.InitDatabase()
		models.InitModels()

		log.Info().Msg("OpenStack Swift Dashboard")
		log.Info().Msg("Scraper")
		log.Info().Msg("")

		targets, err := models.TargetList()
		if err != nil {
			log.Fatal().Err(err)
		}

		scheduler := gocron.NewScheduler(time.UTC)
		for _, target := range targets {
			log.Info().Str("name", target.Name).Str("endpoint", target.Endpoint).Msg("Registering target")
			scheduler.Cron("00,30 * * * *").Do(scrape.ScrapeDisk, target)
		}
		scheduler.StartBlocking()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Output file name and line number.")
	rootCmd.PersistentFlags().StringVar(&dsn, "dsn", "", "Data source name.")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetDefault("verbose", false)

	viper.BindPFlags(rootCmd.Flags())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
