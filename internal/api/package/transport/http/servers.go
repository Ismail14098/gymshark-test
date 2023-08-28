package http

import (
	kithttp "github.com/go-kit/kit/transport/http"

	"gymshark-test/internal/api/package/middleware"
	encoders "gymshark-test/pkg/transport/http"
)

type Servers struct {
	GetAllServer              *kithttp.Server
	GetAllOrderedBySizeServer *kithttp.Server
	StorePackageServer        *kithttp.Server
	UpdatePackageServer       *kithttp.Server
	DeletePackageServer       *kithttp.Server
	PackServer                *kithttp.Server
}

func MakeServers(endpoints *middleware.Endpoints, options []kithttp.ServerOption) *Servers {
	return &Servers{
		GetAllServer: kithttp.NewServer(
			endpoints.GetAllEndpoint,
			emptyDecoder,
			encoders.EncodeResponse,
			options...,
		),
		GetAllOrderedBySizeServer: kithttp.NewServer(
			endpoints.GetAllOrderedBySizeEndpoint,
			emptyDecoder,
			encoders.EncodeResponse,
			options...,
		),
		StorePackageServer: kithttp.NewServer(
			endpoints.StorePackageEndpoint,
			storePackageDecoder,
			encoders.EncodeResponse,
			options...,
		),
		UpdatePackageServer: kithttp.NewServer(
			endpoints.UpdatePackageEndpoint,
			updatePackageDecoder,
			encoders.EncodeResponse,
			options...,
		),
		DeletePackageServer: kithttp.NewServer(
			endpoints.DeletePackageEndpoint,
			deletePackageDecoder,
			encoders.EncodeResponse,
			options...,
		),
		PackServer: kithttp.NewServer(
			endpoints.PackEndpoint,
			packDecoder,
			encoders.EncodeResponse,
			options...,
		),
	}
}
