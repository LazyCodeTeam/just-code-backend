package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/middleware"
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
type currentProfileHandler struct {
	authMiddleware *middleware.AuthTokenValidator
}

func NewCurrentProfileHandler(
	authMiddleware *middleware.AuthTokenValidator,
) Handler {
	return &currentProfileHandler{
		authMiddleware: authMiddleware,
	}
}

func (h *currentProfileHandler) Register(router chi.Router) {
	router.With(h.authMiddleware.Handle).Get("/api/v1/profile/current", h.handleHttp)
}

func (h *currentProfileHandler) handleHttp(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
