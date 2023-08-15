package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

// swagger:route PUT /api/v1/profile/current profile profilePutCurrent
//
// # Update current profile
//
// Responses:
//
//	200: emptyResponse
//	401: errorResponse
//	500: errorResponse
type profilePutCurrentHandler struct {
	updateCurrentprofile *usecase.UpdateCurrentProfile
	validate             *validator.Validate
}

func NewProfilePutCurrentHandler(
	updateCurrentprofile *usecase.UpdateCurrentProfile,
	validate *validator.Validate,
) Handler {
	return &profilePutCurrentHandler{
		updateCurrentprofile: updateCurrentprofile,
		validate:             validate,
	}
}

func (h *profilePutCurrentHandler) Register(router chi.Router) {
	router.Put("/api/v1/profile/current", h.handleHttp)
}

func (h *profilePutCurrentHandler) handleHttp(writer http.ResponseWriter, request *http.Request) {
	body, err := util.DeserializeAndValidateBody[dto.CreateProfileParams](request, h.validate)
	if err != nil {
		util.WriteError(writer, err)

		return
	}
	err = h.updateCurrentprofile.Invoke(request.Context(), dto.CreateProfileParamsToModel(body))
	if err != nil {
		util.WriteError(writer, err)

		return
	}

	writer.WriteHeader(http.StatusOK)
}
