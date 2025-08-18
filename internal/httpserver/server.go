package httpserver

import (
	"generic-backend-service/internal/obs"
	"net/http"
)

type Server struct {
	srv       *http.Server
	router    http.Handler
	readiness *obs.Readiness
}
