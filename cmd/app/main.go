package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpencoders "gymshark-test/pkg/transport/http"
	corsutil "gymshark-test/pkg/utils/cors"
	healthcheckutil "gymshark-test/pkg/utils/healthcheck"
	loggerutil "gymshark-test/pkg/utils/logger"

	"gymshark-test/internal/api/package/middleware"
	package_httptransport "gymshark-test/internal/api/package/transport/http"
	"gymshark-test/internal/config"
	"gymshark-test/internal/constructors"
	"gymshark-test/internal/repository/inmem"
)

func main() {
	// parse flags
	httpPort := flag.String("http.port", ":9000", "HTTP listen address")
	flag.Parse()

	logger := loggerutil.NewServiceLogger("gymshark-test")

	_ = level.Info(logger).Log("msg", "service started")

	// init service configuration
	err := config.InitConfigs()
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// init repository (data access layer)
	mainRepo := inmem.NewStore()

	// init service (business-logic layer)
	svc := constructors.NewPackageService(logger, mainRepo)

	// init endpoints (endpoints layer)
	endpoints := middleware.MakeEndpoints(svc)

	// init HTTP handler (transport layer)
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(httpencoders.EncodeErrorResponse),
	}

	packageHandler := package_httptransport.MakeRoutes(package_httptransport.MakeServers(endpoints, serverOptions))

	// add routes, prometheus and health check handlers
	http.Handle("/package-api/v1/", corsutil.CORS(packageHandler))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/check", healthcheckutil.HealthCheck)

	// init errors chan and ticker
	errs := make(chan error)

	// make chan for syscall
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// init HTTP server
	go func() {
		_ = level.Info(logger).Log("transport", "HTTP", "port", *httpPort)
		errs <- http.ListenAndServe(*httpPort, nil)
	}()

	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
