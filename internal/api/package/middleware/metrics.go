package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"

	"gymshark-test/internal/api/package/service"
	"gymshark-test/internal/api/package/transport"
)

// Instrumenting middleware entity
type metricsMiddleware struct {
	next           service.PackageService
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

// NewInstrumentingMiddleware Instrumenting middleware constructor
func NewInstrumentingMiddleware(counter, counterErr metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next service.PackageService) service.PackageService {
		return &metricsMiddleware{
			next:           next,
			requestCount:   counter,
			requestError:   counterErr,
			requestLatency: latency,
		}
	}
}

// Instrumenting middleware private method
func (im *metricsMiddleware) metrics(begin time.Time, method string, err error) {
	im.requestCount.With("method", method).Add(1)
	if err != nil {
		im.requestError.With("method", method).Add(1)
	}
	im.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

func (im *metricsMiddleware) GetAll(ctx context.Context) (resp *transport.PackagesResponse) {
	defer im.metrics(time.Now(), "GetAll", nil)
	return im.next.GetAll(ctx)
}

func (im *metricsMiddleware) GetAllOrderedBySize(ctx context.Context) (resp *transport.PackagesResponse) {
	defer im.metrics(time.Now(), "GetAllOrderedBySize", nil)
	return im.next.GetAllOrderedBySize(ctx)
}

func (im *metricsMiddleware) StorePackage(ctx context.Context, req *transport.StorePackageRequest) (resp *transport.PackageResponse, err error) {
	defer im.metrics(time.Now(), "StorePackage", err)
	return im.next.StorePackage(ctx, req)
}

func (im *metricsMiddleware) UpdatePackage(ctx context.Context, req *transport.UpdatePackageRequest) (resp *transport.PackageResponse, err error) {
	defer im.metrics(time.Now(), "UpdatePackage", err)
	return im.next.UpdatePackage(ctx, req)
}

func (im *metricsMiddleware) DeletePackage(ctx context.Context, req *transport.DeletePackageRequest) (resp *transport.StatusResponse, err error) {
	defer im.metrics(time.Now(), "DeletePackage", err)
	return im.next.DeletePackage(ctx, req)
}

func (im *metricsMiddleware) Pack(ctx context.Context, req *transport.PackRequest) (resp *transport.PackResponse) {
	defer im.metrics(time.Now(), "DeletePackage", nil)
	return im.next.Pack(ctx, req)
}
