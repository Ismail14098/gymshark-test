package repository

type InMemStore interface {
	Package() PackageRepository
}
