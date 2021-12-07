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
		internal.InitZerolog(true)
		databases.InitDatabase()
		models.InitModels()

		scheduler := gocron.NewScheduler(time.UTC)
		scheduler.Cron("* * * * *").Do(scrape.ScrapeDisk)
		scheduler.StartBlocking()
	},
}

func init() {
	log.Info().Msg("init log")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Output file name and line number.")
	rootCmd.PersistentFlags().StringVar(&dsn, "dsn", "", "Data source name.")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.BindPFlag("dsn", rootCmd.Flags().Lookup("dsn"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
