package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

const dryRunQueryParam = "dry_run"

// swagger:parameters contentPut
type uploadContentParams struct {
	// in: body
	Body []dto.ExpectedTechnology
	// If true, then no changes will be committed to database.
	//
	// in: query
	DryRun bool `json:"dry_run"`
}

type AdminContentHandler struct {
	validate      *validator.Validate
	uploadContent *usecase.UploadContent
}

func NewAdminContentHandler(
	validate *validator.Validate,
	uploadContent *usecase.UploadContent,
) *AdminContentHandler {
	return &AdminContentHandler{
		validate:      validate,
		uploadContent: uploadContent,
	}
}

func (h *AdminContentHandler) Register(router chi.Router) {
	router.Route("/v1/content", func(router chi.Router) {
		// swagger:route GET /admin/api/v1/content admin contentPut
		//
		// Upload content
		//
		// Takes expected state of content and updates state of content in database to match expected state.
		//
		// Responses:
		//   204: noContentResponse
		//   400: errorResponse
		//   401: errorResponse
		//   500: errorResponse
		router.Put("/", h.handlePutUploadContent)
	})
}

func (h *AdminContentHandler) handlePutUploadContent(w http.ResponseWriter, r *http.Request) {
	body, err := util.DeserializeAndValidateBody[[]dto.ExpectedTechnology](r, h.validate)
	dryRun := r.URL.Query().Get(dryRunQueryParam) == "true"
	if err != nil {
		util.WriteError(w, err)

		return
	}
	expected := dto.ExpectedTechnologiesSliceToDomain(body)
	err = h.uploadContent.Invoke(r.Context(), expected, dryRun)
	if err != nil {
		util.WriteError(w, err)

		return
	}
	w.WriteHeader(http.StatusNoContent)
}
