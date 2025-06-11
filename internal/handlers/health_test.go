package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/robkenis/TrustTap/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := handlers.Health()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, rr.Body.String(), `{"status":"UP"}`)
}
