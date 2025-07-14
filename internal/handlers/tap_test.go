package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
		handler := Tap()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})
	t.Run("X-Real-IP header is used", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/tap", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-Real-IP", "0.0.0.2")

		rr := httptest.NewRecorder()
		handler := Tap()
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
		handler := Tap()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})
}
