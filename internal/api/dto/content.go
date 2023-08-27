package dto

import (
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type TaskContentType string

const (
	TaskContentTypeLesson          TaskContentType = "LESSON"
	TaskContentTypeSingleSelection TaskContentType = "SINGLE_SELECTION"
	TaskContentTypeMultiSelection  TaskContentType = "MULTI_SELECTION"
)

type Technology struct {
	// Technology id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Technology name
	//
	// required: true
	// min length: 1
	// max length: 1024
	Name string `json:"name"`
	// Technology description
	//
	// min length: 1
	// required: false
	Description *string `json:"description"`
	// Technology image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url"`
	// Preview of technology sections
	//
	// required: true
	// min length: 1
	Section []SectionPreview `json:"sections"`
}

type SectionPreview struct {
	// Section id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Section name
	//
	// required: true
	// min length: 1
	// max length: 1024
	Name string `json:"name"`
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
