package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cqroot/openstack-swift-dashboard/controllers"
	"github.com/cqroot/openstack-swift-dashboard/databases"
	"github.com/cqroot/openstack-swift-dashboard/internal"
	"github.com/cqroot/openstack-swift-dashboard/models"
)

var (
	verbose bool
	dsn     string
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitZerolog()
		databases.InitDatabase()
		models.InitModels()

		r := gin.Default()

		r.Static("/ui", "./web/dist")

		v1Group := r.Group("/v1")
		initV1Group(v1Group)

		r.Run(":8088")
	},
}

func init() {
	log.Info().Msg("init log")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Output file name and line number.")
	rootCmd.PersistentFlags().StringVar(&dsn, "dsn", "", "Data source name.")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.BindPFlags(rootCmd.Flags())
}

func initV1Group(v1Group *gin.RouterGroup) {
	targetGroup := v1Group.Group("/target")
	targetGroup.GET("", controllers.GetTargetList)
	targetGroup.PUT("", controllers.PutTarget)

	diskGroup := v1Group.Group("/disk")
	diskGroup.GET(":id", controllers.GetDiskList)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
