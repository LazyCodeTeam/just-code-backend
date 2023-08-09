package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

// swagger:route GET /api/v1/profile/current profile currentProfile
//
// # Get current profile
//
// Responses:
//
//	200: emptyResponse
//	500: errorResponse
type profileGetCurrentHandler struct {
	getCurrentUser *usecase.GetCurrentUser
}

func NewProfileGetCurrentHandler(
	getCurrentUser *usecase.GetCurrentUser,
) Handler {
	return &profileGetCurrentHandler{
		getCurrentUser: getCurrentUser,
	}
}

func (h *profileGetCurrentHandler) Register(router chi.Router) {
	router.Get("/api/v1/profile/current", h.handleHttp)
}

func (h *profileGetCurrentHandler) handleHttp(writer http.ResponseWriter, request *http.Request) {
	profile, err := h.getCurrentUser.Invoke(request.Context())
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	util.WriteResponseJson(writer, profile)
}
