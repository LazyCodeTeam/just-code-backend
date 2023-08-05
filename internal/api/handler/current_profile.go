package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// swagger:route GET /api/v1/profile/current profile currentProfile
//
// # Get current profile
//
// Responses:
//
//	200: emptyResponse
//	500: errorResponse
type currentProfileHandler struct{}

func NewCurrentProfileHandler() Handler {
	return &currentProfileHandler{}
}

func (h *currentProfileHandler) Register(router chi.Router) {
	router.Get("/api/v1/profile/current", h.handleHttp)
}

func (h *currentProfileHandler) handleHttp(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
