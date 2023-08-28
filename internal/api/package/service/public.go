package service

import (
	"context"
	"gymshark-test/internal/api/package/transport"
	"gymshark-test/internal/models"
)

func (s *Service) GetAll(ctx context.Context) *transport.PackagesResponse {
	return &transport.PackagesResponse{Packages: s.store.Package().GetAll(ctx)}
}

func (s *Service) GetAllOrderedBySize(ctx context.Context) *transport.PackagesResponse {
	return &transport.PackagesResponse{Packages: s.store.Package().GetAllOrderedBySize(ctx)}
}

func (s *Service) StorePackage(ctx context.Context, req *transport.StorePackageRequest) (*transport.PackageResponse, error) {
	pkg, err := s.store.Package().Store(ctx, req.Size)
	if err != nil {
		return nil, err
	}

	return &transport.PackageResponse{Package: pkg}, nil
}

func (s *Service) UpdatePackage(ctx context.Context, req *transport.UpdatePackageRequest) (*transport.PackageResponse, error) {
	pkg := req.ToPackage()

	if err := s.store.Package().Update(ctx, pkg); err != nil {
		return nil, err
	}

	return &transport.PackageResponse{Package: pkg}, nil
}

func (s *Service) DeletePackage(ctx context.Context, req *transport.DeletePackageRequest) (*transport.StatusResponse, error) {
	if err := s.store.Package().Delete(ctx, req.ID); err != nil {
		return nil, err
	}

	return &transport.StatusResponse{Success: true}, nil
}

func (s *Service) Pack(ctx context.Context, req *transport.PackRequest) *transport.PackResponse {
	packages := s.store.Package().GetAllOrderedBySize(ctx)
	var result []*models.PackResult
	for pkg, amount := range calculatePacks(packages, req.ItemsOrdered) {
		result = append(result, &models.PackResult{
			Package: pkg,
			Amount:  amount,
		})
	}
	return &transport.PackResponse{Result: result}
}
