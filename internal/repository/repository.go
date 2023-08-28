package repository

import (
	"context"
	"gymshark-test/internal/models"
)

type (
	PackageRepository interface {
		GetAll(ctx context.Context) []*models.Package
		GetAllOrderedBySize(ctx context.Context) []*models.Package
		Store(ctx context.Context, size uint) (*models.Package, error)
		Update(ctx context.Context, pkg *models.Package) error
		Delete(ctx context.Context, id uint) error
	}
)
