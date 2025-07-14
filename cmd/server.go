package main

import (
	"net/http"
	"os"
	"time"

	"github.com/robkenis/TrustTap/internal/handlers"
	"github.com/robkenis/TrustTap/internal/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const port = "8080"

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}).With().Timestamp().Logger()

	r := http.NewServeMux()

	webDirectory := os.Getenv("WEB_DIR")
	if webDirectory == "" {
		webDirectory = "web"
	}

	r.Handle("GET /health", handlers.Health())
	r.Handle("GET /", http.FileServer(http.Dir(webDirectory)))
	r.Handle("POST /tap", handlers.NewTapHandler(storage.NewInMemoryStorage()))

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	log.Info().Str("port", port).Msg("Starting server...")
	log.Fatal().Err(srv.ListenAndServe())
}
