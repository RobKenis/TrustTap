package main

import (
	"net/http"
	"os"
	"time"

	"github.com/robkenis/TrustTap/internal/components"
	"github.com/robkenis/TrustTap/internal/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const port = "8080"

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}).With().Timestamp().Logger()

	r := http.NewServeMux()

	r.Handle("GET /health", handlers.Health())
	r.Handle("GET /", handlers.Component(components.Tap()))
	r.Handle("POST /tap", handlers.Tap())

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	log.Info().Str("port", port).Msg("Starting server...")
	log.Fatal().Err(srv.ListenAndServe())
}
