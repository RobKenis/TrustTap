package storage

import "github.com/robkenis/TrustTap/internal/model"

type InMemoryStorage struct {
	requests []model.AccessRequest
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		requests: make([]model.AccessRequest, 0),
	}
}

func (s *InMemoryStorage) Store(req model.AccessRequest) error {
	s.requests = append(s.requests, req)
	return nil
}

func (s *InMemoryStorage) All() ([]model.AccessRequest, error) {
	return s.requests, nil
}
