package main

import (
	"fmt"
	"generic-backend-service/internal/config"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello World")
	cfg := config.GetDefault()
	ConfigureLogger(cfg.LoggerInfo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Getting health")
		w.WriteHeader(200)
	})
	mux.HandleFunc("GET /metrics", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Getting metrics")
		w.WriteHeader(200)
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Getting readyz")
		w.WriteHeader(200)
	})

	if err := http.ListenAndServe(cfg.HttpInfo.HTTPAddr, mux); err != nil {
		log.Error().Err(err).Msg("Http server could not be made")
		os.Exit(1)
	}

}

func ConfigureLogger(info *config.LoggerInfo) {
	zerolog.SetGlobalLevel(info.LogLevel)
	log.Log().Msg("finished configuring logger")
}
