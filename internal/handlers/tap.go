package handlers

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/robkenis/TrustTap/internal/model"
	"github.com/robkenis/TrustTap/internal/storage"
	"github.com/rs/zerolog/log"
)

type TapHandler struct {
	storage storage.RequestStorage
}

func NewTapHandler(storage storage.RequestStorage) *TapHandler {
	return &TapHandler{
		storage: storage,
	}
}

type tapResponse struct {
	IpAddress string `json:"ip_address"`
	Status    string `json:"status"`
}

func (h *TapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip, _ = extractHost(r.RemoteAddr)
	}

	request := model.NewAccessRequest(ip)
	err := h.storage.Store(request)
	if err != nil {
		log.Error().Err(err).Msg("Failed to store access request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(tapResponse{
		IpAddress: request.IpAddress,
		Status:    string(request.State),
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
