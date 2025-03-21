package goindexer

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-logr/logr"
	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/pkg/logging"
	"github.com/v1gn35h7/goshell/pkg/service"
)

func NewStartCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "goindexer start",
		Long:  "Starts goindexer cli process",
		Run: func(cmd *cobra.Command, args []string) {
			// Set-up logger
			logger := logging.Logger()
			logger.Info("Logger initated...")

			// Metrics setup
			fieldkeys := []string{"topic", "partition", "thread", "error"}
			consumerMetrics := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
				Namespace: "goshell",
				Subsystem: "goIndexer",
				Name:      "processed_mesages",
				Help:      "Number of requests received.",
			}, fieldkeys)

			var g group.Group
			{
				// Http handlers
				r := http.NewServeMux()
				r.Handle("/metrics", promhttp.Handler())
				r.Handle("/debug/pprof", http.DefaultServeMux)

				g.Add(func() error {
					return http.ListenAndServe(":8080", r)
				}, func(er error) {
					//
				})
			}
			{
				g.Add(func() error {
					// start service
					bootStrapServer(logger, consumerMetrics)
					return nil
				}, func(err error) {

				})
			}

			logger.Info("exit", g.Run())

		},
	}
	return startCmd
}

func bootStrapServer(logger logr.Logger, consumerMetrics metrics.Counter) {

	viper.SetConfigName("app")      // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath) // path to look for the config file in
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file

	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println(viper.AllSettings())

	service.IndexerService(logger, consumerMetrics)
}
