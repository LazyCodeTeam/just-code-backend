package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
)

// swagger:route PUT /api/v1/profile/current profile currentProfile
//
// # Get current profile
//
// Responses:
//
//	200: emptyResponse
//	500: errorResponse
type profilePutCurrentHandler struct {
	db *db.Queries
}

func NewProfilePutCurrentHandler(db *db.Queries) Handler {
	return &profilePutCurrentHandler{
		db,
	}
}

func (h *profilePutCurrentHandler) Register(router chi.Router) {
	router.Put("/api/v1/profile/current", h.handleHttp)
}

func (h *profilePutCurrentHandler) handleHttp(writer http.ResponseWriter, request *http.Request) {
	_, err := h.db.CreateProfile(request.Context(), db.CreateProfileParams{
		ID:   "1",
		Name: "test",
	})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
