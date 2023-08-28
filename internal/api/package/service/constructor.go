package service

import (
	"github.com/go-kit/kit/log"
	"gymshark-test/internal/repository"
)

type Service struct {
	log   log.Logger
	store repository.InMemStore
}

func NewService(log log.Logger, store repository.InMemStore) PackageService {
	return &Service{
		log:   log,
		store: store,
	}
}
