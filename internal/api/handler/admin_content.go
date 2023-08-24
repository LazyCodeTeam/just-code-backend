package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

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
		router.Put("/", h.handlePutUploadTasks)
		router.Put("/dry-run", h.handlePutUploadTasksDryRun)
	})
}

func (h *AdminContentHandler) handlePutUploadTasks(w http.ResponseWriter, r *http.Request) {
	body, err := util.DeserializeAndValidateBody[[]dto.ExpectedTechnology](r, h.validate)
	if err != nil {
		util.WriteError(w, err)

		return
	}
	expected := dto.ExpectedTechnologiesSliceToDomain(body)
	err = h.uploadContent.Invoke(r.Context(), expected)
	if err != nil {
		util.WriteError(w, err)

		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *AdminContentHandler) handlePutUploadTasksDryRun(w http.ResponseWriter, r *http.Request) {
}
