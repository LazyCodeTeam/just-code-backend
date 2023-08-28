package dto

import (
	"time"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type TaskContentType string

const (
	TaskContentTypeLesson          TaskContentType = "LESSON"
	TaskContentTypeSingleSelection TaskContentType = "SINGLE_SELECTION"
	TaskContentTypeMultiSelection  TaskContentType = "MULTI_SELECTION"
)

// TechnologyDto
//
// Represents technology with preview of sections.
//
// swagger:model
type Technology struct {
	// Technology id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Technology name
	//
	// required: true
	Name string `json:"name"`
	// Technology description
	//
	// required: false
	Description *string `json:"description"`
	// Technology image url
	//
	// required: false
	ImageUrl *string `json:"image_url"`
	// Preview of technology sections
	//
	// required: true
	Section []SectionPreview `json:"sections"`
}

// SectionPreviewDto
//
// Represents section preview.
//
// swagger:model
type SectionPreview struct {
	// Section id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Section name
	//
	// required: true
	Name string `json:"name"`
}

// SectionDto
//
// Represents section with preview of tasks.
//
// swagger:model
type Section struct {
	// Section id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"          validate:"required,uuid"`
	// Section name
	//
	// required: true
	Name string `json:"name"        validate:"required,min=1,max=1024"`
	// Section description
	//
	// required: false
	Description *string `json:"description" validate:"omitempty,min=1"`
	// Section image url
	//
	// required: false
	ImageUrl *string `json:"image_url"   validate:"omitempty,min=1"`
	// Preview of section tasks
	//
	// required: true
	Tasks []TaskPreview `json:"tasks"       validate:"required,min=1,dive"`
}

// TaskPreviewDto
//
// Represents task preview.
//
// swagger:model
type TaskPreview struct {
	// Section id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Section name
	//
	// required: true
	Name string `json:"name"`
	// Is task public - not public task will not be accessible for anonymous users
	//
	// required: true
	IsPublic bool `json:"is_public"`
	// Date when task was done - nil if task was not done
	//
	// required: false
	DoneAt *time.Time `json:"done_at,omitempty"`
}

func TechnologyFromDomain(technology model.TechnologyWithSectionsPreview) Technology {
	return Technology{
		Id:          technology.Id,
		Name:        technology.Title,
		Description: technology.Description,
		ImageUrl:    technology.ImageUrl,
		Section:     util.MapSlice(technology.Sections, sectionPreviewFromDomain),
	}
}

func sectionPreviewFromDomain(section model.SectionPreview) SectionPreview {
	return SectionPreview{
		Id:   section.Id,
		Name: section.Title,
	}
}

func SectionFromDomain(section model.SectionWithTasksPreview) Section {
	return Section{
		Id:          section.Id,
		Name:        section.Title,
		Description: section.Description,
		ImageUrl:    section.ImageUrl,
		Tasks:       util.MapSlice(section.Tasks, taskPreviewFromDomain),
	}
}

func taskPreviewFromDomain(task model.TaskPreview) TaskPreview {
	return TaskPreview{
		Id:       task.Id,
		Name:     task.Title,
		IsPublic: task.IsPublic,
		DoneAt:   task.DoneAt,
	}
}
