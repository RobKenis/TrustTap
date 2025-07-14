package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/robkenis/TrustTap/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestTap(t *testing.T) {
	t.Run("X-Forwarded-For header is used", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/tap", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-Forwarded-For", "0.0.0.1")

		rr := httptest.NewRecorder()
		storage := storage.NewInMemoryStorage()
		handler := NewTapHandler(storage)
		handler.ServeHTTP(rr, req)

		requests, _ := storage.All()

		assert.Equal(t, http.StatusAccepted, rr.Code)
		assert.Len(t, requests, 1, "Expected one request to be stored")

		assert.Equal(t, "0.0.0.1", requests[0].IpAddress, "Expected IP to match")
	})
	t.Run("X-Real-IP header is used", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/tap", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-Real-IP", "0.0.0.2")

		rr := httptest.NewRecorder()
		handler := NewTapHandler(storage.NewInMemoryStorage())
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})
	t.Run("RemoteAddr is used", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/tap", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.RemoteAddr = "0.0.0.3"

		rr := httptest.NewRecorder()
		handler := NewTapHandler(storage.NewInMemoryStorage())
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})
}
