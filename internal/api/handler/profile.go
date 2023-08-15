package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

// swagger:response profileGetCurrentResponse
type profileGetCurrentResponse struct {
	// The error message
	// in: body
	Body dto.Profile
}

type profileHandler struct {
	getCurrentUser       *usecase.GetCurrentUser
	updateCurrentprofile *usecase.UpdateCurrentProfile
	validate             *validator.Validate
}

func NewProfileHandler(
	getCurrentUser *usecase.GetCurrentUser,
	updateCurrentprofile *usecase.UpdateCurrentProfile,
	validate *validator.Validate,
) Handler {
	return &profileHandler{
		getCurrentUser:       getCurrentUser,
		updateCurrentprofile: updateCurrentprofile,
		validate:             validate,
	}
}

func (h *profileHandler) Register(router chi.Router) {
	router.Route("/v1/profile", func(router chi.Router) {
		// swagger:route GET /api/v1/profile/current profile profileGetCurrent
		//
		// # Get current profile
		//
		// Returns current profile. If profile does not exist 404 error will be returned.
		//
		// Responses:
		//
		//	200: profileGetCurrentResponse
		//	401: errorResponse
		//	404: errorResponse
		//	500: errorResponse
		router.Get("/current", h.handleGetCurrent)
		// swagger:route PUT /api/v1/profile/current profile profilePutCurrent
		//
		// # Update current profile
		//
		// Creates new profile for current user or updates existing one. If profile already exists all fields will be updated.
		// Nickname must be unique, otherwise 409 error will be returned.
		//
		// Responses:
		//
		//	200: emptyResponse
		//	401: errorResponse
		//	409: errorResponse
		//	500: errorResponse
		router.Put("/current", h.handlePutCurrent)
	})
}

func (h *profileHandler) handleGetCurrent(writer http.ResponseWriter, request *http.Request) {
	profile, err := h.getCurrentUser.Invoke(request.Context())
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	util.WriteResponseJson(writer, dto.ProfileFromModel(profile))
}

func (h *profileHandler) handlePutCurrent(writer http.ResponseWriter, request *http.Request) {
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
