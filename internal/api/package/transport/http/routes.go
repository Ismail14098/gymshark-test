package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRoutes(servers *Servers) *mux.Router {
	router := mux.NewRouter()

	// swagger:route GET /package-api/v1/package/ Package GetAllRequest
	// Package get
	// responses:
	//   200: PackagesResponse
	router.Path("/package-api/v1/package/").
		Methods(http.MethodGet).
		Handler(servers.GetAllServer)

	// swagger:route GET /package-api/v1/package/ordered-by-size/ Package GetAllOrderedBySizeRequest
	// Package get ordered by size
	// responses:
	//   200: PackagesResponse
	router.Path("/package-api/v1/package/ordered-by-size/").
		Methods(http.MethodGet).
		Handler(servers.GetAllOrderedBySizeServer)

	// swagger:route POST /package-api/v1/package/ Package StorePackageRequest
	// Package create
	// responses:
	//   200: PackageResponse
	router.Path("/package-api/v1/package/").
		Methods(http.MethodPost).
		Handler(servers.StorePackageServer)

	// swagger:route PUT /package-api/v1/package/{id}/ Package UpdatePackageRequest
	// Package update
	// responses:
	//   200: PackageResponse
	router.Path("/package-api/v1/package/{id}/").
		Methods(http.MethodPut).
		Handler(servers.UpdatePackageServer)

	// swagger:route DELETE /package-api/v1/package/{id}/ Package DeletePackageRequest
	// Package delete
	// responses:
	//   200: StatusResponse
	router.Path("/package-api/v1/package/{id}/").
		Methods(http.MethodDelete).
		Handler(servers.DeletePackageServer)

	// swagger:route POST /package-api/v1/pack/ Pack PackRequest
	// Package create
	// responses:
	//   200: PackResponse
	router.Path("/package-api/v1/pack/").
		Methods(http.MethodPost).
		Handler(servers.PackServer)

	return router
}
