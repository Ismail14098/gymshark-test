package middleware

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"gymshark-test/internal/api/package/service"
	"gymshark-test/internal/api/package/transport"
)

type Endpoints struct {
	GetAllEndpoint              endpoint.Endpoint
	GetAllOrderedBySizeEndpoint endpoint.Endpoint
	StorePackageEndpoint        endpoint.Endpoint
	UpdatePackageEndpoint       endpoint.Endpoint
	DeletePackageEndpoint       endpoint.Endpoint
	PackEndpoint                endpoint.Endpoint
}

func MakeEndpoints(s service.PackageService) *Endpoints {
	return &Endpoints{
		GetAllEndpoint:              makeGetAllEndpoint(s),
		GetAllOrderedBySizeEndpoint: makeGetAllOrderedBySizeEndpoint(s),
		StorePackageEndpoint:        makeStorePackageEndpoint(s),
		UpdatePackageEndpoint:       makeUpdatePackageEndpoint(s),
		DeletePackageEndpoint:       makeDeletePackageEndpoint(s),
		PackEndpoint:                makePackEndpoint(s),
	}
}

func makeGetAllEndpoint(s service.PackageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.GetAll(ctx), nil
	}
}

func makeGetAllOrderedBySizeEndpoint(s service.PackageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.GetAllOrderedBySize(ctx), nil
	}
}

func makeStorePackageEndpoint(s service.PackageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.StorePackage(ctx, request.(*transport.StorePackageRequest))
	}
}

func makeUpdatePackageEndpoint(s service.PackageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.UpdatePackage(ctx, request.(*transport.UpdatePackageRequest))
	}
}

func makeDeletePackageEndpoint(s service.PackageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.DeletePackage(ctx, request.(*transport.DeletePackageRequest))
	}
}

func makePackEndpoint(s service.PackageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.Pack(ctx, request.(*transport.PackRequest)), nil
	}
}
