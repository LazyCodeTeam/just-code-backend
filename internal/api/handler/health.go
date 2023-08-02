package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// swagger:route GET /api/health health
//
// # Health check
//
// This will check if the service is up and running.
//
// Responses:
//
//	200: emptyResponse
//	500: errorResponse
type healthHandler struct{}

func NewHealthHandler() Handler {
	return &healthHandler{}
}

func (h *healthHandler) Register(router chi.Router) {
	router.Get("/api/health", h.handleHttp)
}

func (h *healthHandler) handleHttp(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
