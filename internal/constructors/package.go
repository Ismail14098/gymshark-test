package constructors

import (
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"gymshark-test/internal/api/package/middleware"
	"gymshark-test/internal/api/package/service"
	"gymshark-test/internal/repository"
)

func NewPackageService(log log.Logger, store repository.InMemStore) service.PackageService {
	svc := service.NewService(log, store)
	svc = middleware.NewLoggingMiddleware(log)(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "test_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "test_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "test_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}))(svc)
	return svc
}
