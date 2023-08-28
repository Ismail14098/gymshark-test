package service

import (
	"context"

	"gymshark-test/internal/api/package/transport"
)

type PackageService interface {
	GetAll(ctx context.Context) *transport.PackagesResponse
	GetAllOrderedBySize(ctx context.Context) *transport.PackagesResponse
	StorePackage(ctx context.Context, req *transport.StorePackageRequest) (*transport.PackageResponse, error)
	UpdatePackage(ctx context.Context, req *transport.UpdatePackageRequest) (*transport.PackageResponse, error)
	DeletePackage(ctx context.Context, req *transport.DeletePackageRequest) (*transport.StatusResponse, error)
	Pack(ctx context.Context, req *transport.PackRequest) *transport.PackResponse
}
