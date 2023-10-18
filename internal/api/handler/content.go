package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
	coreUtil "github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

// swagger:response contentGetTechnologiesResponse
type contentGetTechnologiesResponse struct {
	// in: body
	Body []dto.Technology
}

// swagger:response contentGetSectionsResponse
type getSectionsResponse struct {
	// in: body
	Body []dto.Section
}

// swagger:response contentGetTasksResponse
type contentGetTasksResponse struct {
	// in: body
	Body []dto.Task
}

// swagger:parameters sectionsGet
type getSectionsArgs struct {
	// in: path
	TechnologyId string `json:"technologyId"`
}

// swagger:parameters tasksGet
type getTasksArgs struct {
	// in: path
	SectionId string `json:"sectionId"`
}

type contentHandler struct {
	validate        *validator.Validate
	getTechnologies *usecase.GetTechnologies
	getSections     *usecase.GetSections
	getTasks        *usecase.GetTasks
}

func NewContentHandler(
	validate *validator.Validate,
	getTechnologies *usecase.GetTechnologies,
	getSections *usecase.GetSections,
	getTasks *usecase.GetTasks,
) *contentHandler {
	return &contentHandler{
		validate:        validate,
		getTechnologies: getTechnologies,
		getSections:     getSections,
		getTasks:        getTasks,
	}
}

func (h *contentHandler) Register(router chi.Router) {
	router.Route("/v1/content", func(r chi.Router) {
		// swagger:route GET /api/v1/content/technologies content technologiesGet
		//
		// Get technologies
		//
		// This will return all technologies along with preview of their sections
		//
		// Responses:
		//   200: contentGetTechnologiesResponse
		//   401: errorResponse
		//   500: errorResponse
		r.Get("/technologies", h.handleGetTechnologies)
		// swagger:route GET /api/v1/content/technology/{technologyId}/sections content sectionsGet
		//
		// Get technology sections
		//
		// This will return all sections of technology with preview of their tasks
		//
		// Responses:
		//   200: contentGetSectionsResponse
		//   401: errorResponse
		//   500: errorResponse
		r.Get("/technology/{technologyId}/sections", h.handleGetSections)
		// swagger:route GET /api/v1/content/section/{sectionId}/tasks content tasksGet
		//
		// Get section tasks
		//
		// This will return all tasks of section
		//
		// Responses:
		//   200: contentGetTasksResponse
		//   401: errorResponse
		//   500: errorResponse
		r.Get("/section/{sectionId}/tasks", h.handleGetTasks)
	})
}

func (h *contentHandler) handleGetTechnologies(w http.ResponseWriter, r *http.Request) {
	technologies, err := h.getTechnologies.Invoke(r.Context())
	if err != nil {
		util.WriteError(w, err)
		return
	}

	dtos := coreUtil.MapSlice[model.TechnologyWithSectionsPreview, dto.Technology](
		technologies,
		dto.TechnologyFromDomain,
	)

	util.WriteResponseJson(w, dtos)
}

func (h *contentHandler) handleGetSections(w http.ResponseWriter, r *http.Request) {
	technologyId := chi.URLParam(r, "technologyId")

	sections, err := h.getSections.Invoke(r.Context(), technologyId)
	if err != nil {
		util.WriteError(w, err)
		return
	}
	dtos := coreUtil.MapSlice[model.SectionWithTasksPreview, dto.Section](
		sections,
		dto.SectionFromDomain,
	)

	util.WriteResponseJson(w, dtos)
}

func (h *contentHandler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	sectionId := chi.URLParam(r, "sectionId")

	tasks, err := h.getTasks.Invoke(r.Context(), sectionId)
	if err != nil {
		util.WriteError(w, err)
		return
	}

	dtos := coreUtil.MapSlice[model.Task, dto.Task](
		tasks,
		dto.TaskFromDomain,
	)

	util.WriteResponseJson(w, dtos)
}
