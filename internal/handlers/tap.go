package handlers

import (
	"net"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Tap() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.Header.Get("X-Real-IP")
		}
		if ip == "" {
			ip, _ = extractHost(r.RemoteAddr)
		}

		log.Info().Str("ip", ip).Msg("Incoming request")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})
}

// Helper to extract host and port
func extractHost(addr string) (string, error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return addr, err
	}
	return host, nil
}
