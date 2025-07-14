package storage

import (
	"slices"

	"github.com/robkenis/TrustTap/internal/model"
	"github.com/rs/zerolog/log"
)

type InMemoryStorage struct {
	requests []model.AccessRequest
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		requests: make([]model.AccessRequest, 0),
	}
}

func (s *InMemoryStorage) Store(req model.AccessRequest) error {
	if s.shouldStore(req) {
		s.requests = append(s.requests, req)
		log.Info().Str("ip", req.IpAddress).Msg("Access request stored")
	} else {
		log.Info().Str("ip", req.IpAddress).Msg("Duplicate access request ignored")
	}
	return nil
}

func (s *InMemoryStorage) All() ([]model.AccessRequest, error) {
	return s.requests, nil
}

func (s *InMemoryStorage) shouldStore(req model.AccessRequest) bool {
	addressWasRequestedBefore := slices.ContainsFunc(s.requests, func(r model.AccessRequest) bool {
		return r.IpAddress == req.IpAddress
	})
	return !addressWasRequestedBefore
}
