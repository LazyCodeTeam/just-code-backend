package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	app_middleware "github.com/LazyCodeTeam/just-code-backend/internal/api/middleware"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

// swagger:response profileGetCurrentResponse
type profileGetCurrentResponse struct {
	// The error message
	// in: body
	Body dto.Profile
}

// swagger:response profilePutCurrentAvatar
type profilePutCurrentAvatar struct {
	// in: body
	Body []byte
}

type profileHandler struct {
	getCurrentUser       *usecase.GetCurrentUser
	updateCurrentprofile *usecase.UpdateCurrentProfile
	uploadProfileAvatar  *usecase.UploadProfileAvatar
	validate             *validator.Validate
}

func NewProfileHandler(
	getCurrentUser *usecase.GetCurrentUser,
	updateCurrentprofile *usecase.UpdateCurrentProfile,
	uploadProfileAvatar *usecase.UploadProfileAvatar,
	validate *validator.Validate,
) Handler {
	return &profileHandler{
		getCurrentUser:       getCurrentUser,
		updateCurrentprofile: updateCurrentprofile,
		uploadProfileAvatar:  uploadProfileAvatar,
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
		// swagger:route PUT /api/v1/profile/current/avatar profile profilePutCurrentAvatar
		//
		// # Upload current profile avatar.
		//
		// Image must be in jpeg or png format. Max size is 2MB.
		// Should be send as binary data in request body.
		//
		//
		// Responses:
		//
		//	200: emptyResponse
		//	401: errorResponse
		//	500: errorResponse
		router.With(middleware.RequestSize(2*1024*1024)).
			With(app_middleware.AcceptedBodyFileTypes(
				"image/jpeg",
				"image/png",
			)).
			Put("/current/avatar", h.handlePutCurrentAvatar)
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

func (h *profileHandler) handlePutCurrentAvatar(
	writer http.ResponseWriter,
	request *http.Request,
) {
	err := h.uploadProfileAvatar.Invoke(request.Context(), request.Body)
	if err != nil {
		util.WriteError(writer, err)
		return
	}
	writer.WriteHeader(http.StatusOK)
}