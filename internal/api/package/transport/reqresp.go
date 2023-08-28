package transport

import "gymshark-test/internal/models"

type PackageResponse struct {
	Package *models.Package `json:"package"`
}

type StatusResponse struct {
	Success bool `json:"success"`
}

type PackagesResponse struct {
	Packages []*models.Package `json:"packages"`
}

type StorePackageRequest struct {
	Size uint `json:"size"`
}

type UpdatePackageRequest struct {
	ID   uint `json:"-"`
	Size uint `json:"size"`
}

func (r *UpdatePackageRequest) ToPackage() *models.Package {
	return &models.Package{
		ID:   r.ID,
		Size: r.Size,
	}
}

type DeletePackageRequest struct {
	ID uint `json:"-"`
}

type PackRequest struct {
	ItemsOrdered uint `json:"items_ordered"`
}

type PackResponse struct {
	Result []*models.PackResult `json:"result"`
}
