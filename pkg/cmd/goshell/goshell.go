package goshell

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kitlog "github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/goshell/internal/config"
	"github.com/v1gn35h7/goshell/internal/datastore/cassdb"
	"github.com/v1gn35h7/goshell/pkg/logging"
	"github.com/v1gn35h7/goshell/server/pb"
	"github.com/v1gn35h7/goshell/server/service"
	grpctransport "github.com/v1gn35h7/goshell/server/transport/grpc"
	httptransport "github.com/v1gn35h7/goshell/server/transport/http"
	"google.golang.org/grpc"
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
	fmt.Println("        \"\"\"\"\"\"\"                   \"\"\"\"\"\"\"\"\"\"\"\"                ")
	fmt.Println("       \"\"       \"\"                             \"\"    							")
	fmt.Println("      \"\"                                         \"\"    ")
	fmt.Println("      \"\"                                         \"\"     ")
	fmt.Println("      \"\"                \"\"\"\"\"\"             \"\"     ")
	fmt.Println("      \"\"                \"        \"             \"\"     ")
	fmt.Println("      \"\"         \"\"   \"        \"             \"\"     ")
	fmt.Println("       \"\"        \"\"   \"        \"             \"\"     ")
	fmt.Println("         \"\"      \"\"   \"        \"             \"\"     ")
	fmt.Println("           \"\"\"\"\"\"   \"\"\"\"\"\"             \"\"    ")
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
	srvc := service.New(logger)
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

	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		//httpListener, err := net.Listen("tcp", *httpAddr)
		// if err != nil {
		// 	logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		// 	os.Exit(1)
		// }
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", "localhost:8080")
			return srv.ListenAndServe()
		}, func(error) {
			//httpListener.Close()
		})
	}
	{
		// Start gRPC server
		grpcServer := grpctransport.NewGRPCServer(grpctransport.MakeGrpcEndpoints(srvc, logger))
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", "localhost:8082")
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}

		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", "localhost:8082")
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterShellServiceServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())

}
