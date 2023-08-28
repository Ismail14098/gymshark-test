package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"gymshark-test/internal/api/package/service"
	"gymshark-test/internal/api/package/transport"
)

// Logging middleware entity
type loggingMiddleware struct {
	logger log.Logger
	next   service.PackageService
}

// Logging middleware private method
func (lm *loggingMiddleware) logging(begin time.Time, method string, err error) {
	_ = lm.logger.Log("method", method, "took", time.Since(begin), "err", err)
}

// Logging middleware constructor
func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.PackageService) service.PackageService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (lm *loggingMiddleware) GetAll(ctx context.Context) (req *transport.PackagesResponse) {
	defer lm.logging(time.Now(), "GetAll", nil)
	return lm.next.GetAll(ctx)
}

func (lm *loggingMiddleware) GetAllOrderedBySize(ctx context.Context) (resp *transport.PackagesResponse) {
	defer lm.logging(time.Now(), "GetAllOrderedBySize", nil)
	return lm.next.GetAllOrderedBySize(ctx)
}

func (lm *loggingMiddleware) StorePackage(ctx context.Context, req *transport.StorePackageRequest) (resp *transport.PackageResponse, err error) {
	defer lm.logging(time.Now(), "StorePackage", err)
	return lm.next.StorePackage(ctx, req)
}

func (lm *loggingMiddleware) UpdatePackage(ctx context.Context, req *transport.UpdatePackageRequest) (resp *transport.PackageResponse, err error) {
	defer lm.logging(time.Now(), "StorePackage", err)
	return lm.next.UpdatePackage(ctx, req)
}

func (lm *loggingMiddleware) DeletePackage(ctx context.Context, req *transport.DeletePackageRequest) (resp *transport.StatusResponse, err error) {
	defer lm.logging(time.Now(), "StorePackage", err)
	return lm.next.DeletePackage(ctx, req)
}

func (lm *loggingMiddleware) Pack(ctx context.Context, req *transport.PackRequest) (resp *transport.PackResponse) {
	defer lm.logging(time.Now(), "Pack", nil)
	return lm.next.Pack(ctx, req)
}
