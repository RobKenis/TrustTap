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

	t.Run("Prevent duplicate requests", func(t *testing.T) {
		req := model.NewAccessRequest("0.0.0.2")
		err := storage.Store(req)
		assert.NoError(t, err, "Failed to store request")

		requests, err := storage.All()
		assert.NoError(t, err, "Failed to retrieve requests")

		assert.Len(t, requests, 2)
		assert.Equal(t, "0.0.0.1", requests[0].IpAddress)
		assert.Equal(t, "0.0.0.2", requests[1].IpAddress)

		// Try to store the same request again
		err = storage.Store(req)

		assert.NoError(t, err, "Storing a duplicate request should not return an error")
		requests, err = storage.All()

		assert.NoError(t, err, "Failed to retrieve requests after duplicate store attempt")
		assert.Len(t, requests, 2, "Duplicate request should not increase the count")
	})
}
