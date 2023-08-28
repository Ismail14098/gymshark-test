package _package

import t "gymshark-test/internal/api/package/transport"

// swagger:parameters GetAllRequest
type GetAllRequest struct{}

// swagger:parameters GetAllOrderedBySizeRequest
type GetAllOrderedBySizeRequest struct{}

// swagger:response PackageResponse
type PackageResponse struct {
	// in:body
	Body t.PackageResponse
}

// swagger:response PackagesResponse
type PackagesResponse struct {
	// in:body
	Body t.PackagesResponse
}

// swagger:parameters StorePackageRequest
type StorePackageRequest struct {
	// in:body
	Body t.StorePackageRequest
}

// swagger:parameters UpdatePackageRequest
type UpdatePackageRequest struct {
	// in: path
	ID uint `json:"id"`
	// in:body
	Body t.UpdatePackageRequest
}

// swagger:parameters DeletePackageRequest
type DeletePackageRequest struct {
	// in: path
	ID uint `json:"id"`
}

// swagger:response StatusResponse
type StatusResponse struct {
	// in:body
	Body t.StatusResponse
}

// swagger:parameters PackRequest
type PackRequest struct {
	// in:body
	Body t.PackRequest
}

// swagger:response PackResponse
type PackResponse struct {
	// in:body
	Body t.PackResponse
}
