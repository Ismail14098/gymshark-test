package inmem

import (
	"context"
	"sort"
	"sync"

	"gymshark-test/internal/models"
	"gymshark-test/pkg/errors"
)

var db = map[uint]*models.Package{
	0: {
		ID:   0,
		Size: 250,
	},
	1: {
		ID:   1,
		Size: 500,
	},
	2: {
		ID:   2,
		Size: 1000,
	},
	4: {
		ID:   4,
		Size: 5000,
	},
	3: {
		ID:   3,
		Size: 2000,
	},
}

type PackageRepository struct {
	db map[uint]*models.Package
	// existingSizes used in order to guarantee that size will be unique
	existingSizes map[uint]struct{}
	id            uint
	mutex         sync.RWMutex
}

func NewPackageRepository() *PackageRepository {
	existingSizes := make(map[uint]struct{})
	for _, v := range db {
		existingSizes[v.Size] = struct{}{}
	}
	return &PackageRepository{
		db:            db,
		existingSizes: existingSizes,
		id:            4,
		mutex:         sync.RWMutex{},
	}
}

func (p *PackageRepository) GetAll(ctx context.Context) []*models.Package {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	result := make([]*models.Package, len(p.db))
	i := 0
	for _, v := range p.db {
		result[i] = v
		i++
	}
	return result
}

func (p *PackageRepository) GetAllOrderedBySize(ctx context.Context) []*models.Package {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	packages := p.GetAll(ctx)
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Size < packages[j].Size
	})
	return packages
}

func (p *PackageRepository) Store(ctx context.Context, size uint) (*models.Package, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.existingSizes[size]; ok {
		return nil, errors.DBWriteError.SetDevMessage("err package with such size exists")
	}

	id := p.id + 1
	p.id++

	pkg := &models.Package{
		ID:   id,
		Size: size,
	}

	p.db[pkg.ID] = pkg
	p.existingSizes[size] = struct{}{}

	return pkg, nil
}

func (p *PackageRepository) Update(ctx context.Context, pkg *models.Package) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.existingSizes[pkg.Size]; ok {
		return errors.DBWriteError.SetDevMessage("err package with such size exists")
	}

	oldPkg, ok := p.db[pkg.ID]
	if !ok {
		return errors.DBWriteError.SetDevMessage("err package with such id does not exists")
	}

	delete(p.existingSizes, oldPkg.Size)

	p.db[pkg.ID] = pkg
	p.existingSizes[pkg.Size] = struct{}{}

	return nil
}

func (p *PackageRepository) Delete(ctx context.Context, id uint) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	pkg, ok := p.db[id]
	if !ok {
		return errors.DBWriteError.SetDevMessage("err package with such id does not exists")
	}

	delete(p.db, pkg.ID)
	delete(p.existingSizes, pkg.Size)

	return nil
}
