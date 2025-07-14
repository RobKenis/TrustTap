package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockComponent struct{}

func (m mockComponent) Render(ctx context.Context, w io.Writer) error {
	_, _ = w.Write([]byte("<div>mock</div>"))
	return nil
}

func TestComponent_Handler(t *testing.T) {
	comp := mockComponent{}
	h := Component(comp)
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	h.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "expected status 200")
	assert.Equal(t, "text/html", rr.Header().Get("Content-Type"), "expected Content-Type 'text/html'")
	assert.Equal(t, "<div>mock</div>", rr.Body.String(), "expected body '<div>mock</div>'")
}
