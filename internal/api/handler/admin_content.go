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

// swagger:response contentAssetPutResponse
type contentAssetPutResponse struct {
	// in: body
	Body dto.Asset
}

// swagger:parameters contentPutAsset
type contentPutAssetParams struct {
	// in: body
	Body []byte
}

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
	saveAsset     *usecase.SaveAsset
}

func NewAdminContentHandler(
	validate *validator.Validate,
	uploadContent *usecase.UploadContent,
	saveAsset *usecase.SaveAsset,
) *AdminContentHandler {
	return &AdminContentHandler{
		validate:      validate,
		uploadContent: uploadContent,
		saveAsset:     saveAsset,
	}
}

func (h *AdminContentHandler) Register(router chi.Router) {
	router.Route("/v1/content", func(router chi.Router) {
		// swagger:route PUT /admin/api/v1/content admin contentPut
		//
		// Upload content
		//
		// Takes expected state of content and updates state of content in database to match expected state.
		//
		// Responses:
		//   204: emptyResponse
		//   400: errorResponse
		//   401: errorResponse
		//   500: errorResponse
		router.Put("/", h.handlePutContent)
		// swagger:route PUT /admin/api/v1/content/asset admin contentPutAsset
		//
		// Upload content asset
		//
		// Takes asset and uploads it to the storage. Returns url of uploaded asset.
		//
		// Responses:
		//   200: contentAssetPutResponse
		//   400: errorResponse
		//   401: errorResponse
		//   500: errorResponse
		router.Put("/asset", h.handlePutContentAsset)
	})
}

func (h *AdminContentHandler) handlePutContentAsset(w http.ResponseWriter, r *http.Request) {
	asset, err := h.saveAsset.Invoke(r.Context(), r.Body)
	if err != nil {
		util.WriteError(w, err)
		return
	}
	dto := dto.AssetFromDomain(asset)

	util.WriteResponseJson(w, dto, http.StatusCreated)
}

func (h *AdminContentHandler) handlePutContent(w http.ResponseWriter, r *http.Request) {
	body, err := util.DeserializeAndValidateBodySlice[dto.ExpectedTechnology](r, h.validate)
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
