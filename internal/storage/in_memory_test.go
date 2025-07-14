package storage_test

import (
	"testing"

	"github.com/robkenis/TrustTap/internal/model"
	"github.com/robkenis/TrustTap/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryStorage(t *testing.T) {
	storage := storage.NewInMemoryStorage()

	t.Run("Store and retrieve requests", func(t *testing.T) {
		req := model.NewAccessRequest("0.0.0.1")
		err := storage.Store(req)
		if err != nil {
			t.Fatal(err)
		}

		requests, err := storage.All()
		if err != nil {
			t.Fatal(err)
		}

		assert.Len(t, requests, 1)
		assert.Equal(t, "0.0.0.1", requests[0].IpAddress)
	})
}
