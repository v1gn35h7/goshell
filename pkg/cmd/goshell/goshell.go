package goshell

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/internal/config"
	"github.com/v1gn35h7/goshell/internal/datastore/cassdb"
	"github.com/v1gn35h7/goshell/pkg/logging"
	"github.com/v1gn35h7/goshell/server/service"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitlog "github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	httptransport "github.com/v1gn35h7/goshell/server/transport/http"
)

var (
	configPath string
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "goShell Server",
		Short: "goshellctl Server",
		Long:  "Goshell Server started",
		Run: func(cmd *cobra.Command, args []string) {

			//Init
			printLogo()

			//Bootstrap server
			bootStrapServer()
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", "", "config file path")

	return rootCmd
}

func printLogo() {
	fmt.Println("#########################################################")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("        \"\"\"\"\"\"\"                                   ")
	fmt.Println("       \"\"       \"\"                                  ")
	fmt.Println("      \"\"                                              ")
	fmt.Println("      \"\"                                               ")
	fmt.Println("      \"\"                \"\"\"\"\"\"                   ")
	fmt.Println("      \"\"                \"        \"                   ")
	fmt.Println("      \"\"         \"\"   \"        \"                   ")
	fmt.Println("       \"\"        \"\"   \"        \"                   ")
	fmt.Println("         \"\"      \"\"   \"        \"                   ")
	fmt.Println("           \"\"\"\"\"\"   \"\"\"\"\"\"                  ")
	fmt.Println("                                                         ")
	fmt.Println("                                                         ")
	fmt.Println("#########################################################")
}

func bootStrapServer() {
	//Logger setup
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// Init read config
	fmt.Println("Config path set to: ", configPath)
	config.ReadConfig(configPath, logging.Logger())

	// Init database
	cassdb.SetUpSession()

	//Mertics setup
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "goshell",
		Subsystem: "web_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "goshell",
		Subsystem: "web_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// goShell Service init
	srvc := service.New()
	serviceLoggingMiddleware := service.NewLoggingServiceMiddleware(logger, srvc)
	serviceInstrumentationMiddleware := service.NewInstrumentationServiceMiddleware(requestCount, requestLatency, serviceLoggingMiddleware)

	// Mux Routes
	r := httptransport.MakeHandlers(serviceInstrumentationMiddleware, logger)

	port := viper.GetString("goshell.server.port")

	// Start Server
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
