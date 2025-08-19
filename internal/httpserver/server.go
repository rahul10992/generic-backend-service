package httpserver

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type HttpServer struct {
	addr   string
	srv    *http.Server
	router http.Handler
}

type Server interface {
	Shutdown()
}

func GetNewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (h *HttpServer) ListenAndServe() error {
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	if err := http.ListenAndServe(h.addr, mux); err != nil {
		return err
	}

	return nil
}

func (h *HttpServer) Shutdown() {
	// what do I add here?
}

func (h *HttpServer) RegisterRoutes(mux *http.ServeMux) {
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

}
