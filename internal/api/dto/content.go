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
	Id string `json:"id" `
	// Section name
	//
	// required: true
	Name string `json:"name"`
	// Section description
	//
	// required: false
	Description *string `json:"description"`
	// Section image url
	//
	// required: false
	ImageUrl *string `json:"image_url"`
	// Preview of section tasks
	//
	// required: true
	Tasks []TaskPreview `json:"tasks"`
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

// TaskDto
//
// Represents task preview.
//
// swagger:model
type Task struct {
	// Section id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Section name
	//
	// required: true
	Name string `json:"name"`
	// Task description
	//
	// min length: 1
	// required: false
	Description *string `json:"description,omitempty"`
	// Task image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url,omitempty"`
	// Task difficulty
	//
	// required: true
	// min: 1
	// max: 10
	Difficulty int `json:"difficulty" `
	// Date when task was done - nil if task was not done
	//
	// required: false
	DoneAt *time.Time `json:"done_at,omitempty"`
	// Task content -- nil if task is non public and user is anonymous
	//
	// required: false
	Content *TaskContent `json:"content,omitempty"`
}

// TaskContentDto
//
// Represents task content.
//
// swagger:model
type TaskContent struct {
	// Type of content
	//
	// example: LESSON
	Kind TaskContentType `json:"kind"`
	// Task description
	//
	// required: true
	Content string `json:"content"`
	// Possible answers
	//
	// Required if kind is SINGLE_SELECTION or MULTI_SELECTION.
	//
	// required: false
	// min length: 2
	// max length: 64
	Options []TaskOption `json:"options,omitempty"`
	// Index of correct answer
	//
	// Required if kind is SINGLE_SELECTION.
	//
	// required: false
	// min: 0
	// max: 63
	CorrectOption *int `json:"correct_option,omitempty"`
	// Indexes of correct answers
	//
	// Required if kind is MULTI_SELECTION.
	//
	// required: false
	// min length: 1
	// max length: 64
	CorrectOptions []int `json:"correct_options,omitempty"`
	// Task hints
	//
	// required: false
	// max length: 128
	Hints []TaskHint `json:"hints,omitempty"`
}

type TaskOption struct {
	// Option id -- unique only in task
	//
	// required: true
	Id int `json:"id"`
	// Option content
	//
	// required: true
	Content string `json:"content"`
}

type TaskHint struct {
	// Hint id -- unique only in task
	//
	// required: true
	Id int `json:"id"`
	// Hint content
	//
	// required: true
	Content string `json:"content"`
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

func TaskFromDomain(task model.Task) Task {
	return Task{
		Id:          task.Id,
		Name:        task.Title,
		DoneAt:      task.DoneAt,
		Description: task.Description,
		ImageUrl:    task.ImageUrl,
		Difficulty:  task.Difficulty,
		Content:     taskContentFromDomain(task.Content),
	}
}

func taskContentFromDomain(content *model.TaskContent) *TaskContent {
	if content == nil {
		return nil
	}

	switch {
	case content.Lesson != nil:
		return &TaskContent{
			Kind:    TaskContentTypeLesson,
			Content: content.Lesson.Description,
			Hints:   util.MapSlice(content.Lesson.Hints, taskHintFromDomain),
		}
	case content.SingleSelection != nil:
		return &TaskContent{
			Kind:          TaskContentTypeSingleSelection,
			Content:       content.SingleSelection.Description,
			Options:       util.MapSlice(content.SingleSelection.Options, taskOptionFromDomain),
			CorrectOption: &content.SingleSelection.CorrectOptionId,
			Hints:         util.MapSlice(content.SingleSelection.Hints, taskHintFromDomain),
		}
	case content.MultiSelection != nil:
		return &TaskContent{
			Kind:           TaskContentTypeMultiSelection,
			Content:        content.MultiSelection.Description,
			Options:        util.MapSlice(content.MultiSelection.Options, taskOptionFromDomain),
			CorrectOptions: content.MultiSelection.CorrectOptionIds,
			Hints:          util.MapSlice(content.MultiSelection.Hints, taskHintFromDomain),
		}
	}

	return nil
}

func taskHintFromDomain(hint model.Hint) TaskHint {
	return TaskHint{
		Id:      hint.Id,
		Content: hint.Content,
	}
}

func taskOptionFromDomain(option model.Option) TaskOption {
	return TaskOption{
		Id:      option.Id,
		Content: option.Content,
	}
}
