package inmem

import "gymshark-test/internal/repository"

type Store struct {
	PackageRepository repository.PackageRepository
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Package() repository.PackageRepository {
	if s.PackageRepository == nil {
		s.PackageRepository = NewPackageRepository()
	}

	return s.PackageRepository
}
