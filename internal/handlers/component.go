package handlers

import (
	"net/http"

	"github.com/a-h/templ"
)

func Component(comp templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		_ = comp.Render(r.Context(), w)
	})
}
