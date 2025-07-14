package storage

import "github.com/robkenis/TrustTap/internal/model"

type RequestStorage interface {
	Store(request model.AccessRequest) error
	All() ([]model.AccessRequest, error)
}
