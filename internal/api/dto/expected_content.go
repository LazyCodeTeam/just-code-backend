package dto

import "github.com/LazyCodeTeam/just-code-backend/internal/core/model"

// ExpectedTechnologyDto
//
// Uepresents desired state of technology.
// Will be compared with actual state and required changes will be applied
//
// swagger:model
type ExpectedTechnology struct {
	// Technology id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"          validate:"required,uuid"`
	// Technology name
	//
	// required: true
	// min length: 1
	// max length: 1024
	Name string `json:"name"        validate:"required,min=1,max=1024"`
	// Technology description
	//
	// min length: 1
	// required: false
	Description *string `json:"description" validate:"omitempty,min=1"`
	// Technology image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url"   validate:"omitempty,min=1"`
	// Expended sections
	//
	// required: true
	// min length: 1
	ExpectedSections []ExpectedSection `json:"sections"    validate:"required,min=1,dive"`
}

// ExpectedSectionDto
//
// Represents desired state of section.
// Will be compared with actual state and required changes will be applied
//
// swagger:model
type ExpectedSection struct {
	// Section id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"          validate:"required,uuid"`
	// Section name
	//
	// required: true
	// min length: 1
	// max length: 1024
	Name string `json:"name"        validate:"required,min=1,max=1024"`
	// Section description
	//
	// min length: 1
	// required: false
	Description *string `json:"description" validate:"omitempty,min=1"`
	// Section image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url"   validate:"omitempty,min=1"`
	// Expended tasks
	//
	// required: true
	// min length: 1
	Tasks []ExpectedTask `json:"tasks"       validate:"required,min=1,dive"`
}

// ExpectedTaskDto
//
// Represents desired state of task in section.
// Will be compared with actual state and required changes will be applied
//
// swagger:model
type ExpectedTask struct {
	// Task id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"          validate:"required,uuid"`
	// Task title
	//
	// required: true
	// min length: 1
	// max length: 1024
	Name string `json:"name"        validate:"required,min=1,max=1024"`
	// Task description
	//
	// min length: 1
	// required: false
	Description *string `json:"description" validate:"omitempty,min=1"`
	// Task image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url"   validate:"omitempty,min=1"`
	// Task difficulty
	//
	// required: true
	// min: 1
	// max: 10
	Difficulty int `json:"difficulty"  validate:"required,min=1,max=10"`
	// Task is public - content will be visible aslo for anonymous users
	//
	// required: true
	IsPublic bool `json:"is_public"`
	// Task is dynamic - content will not be displayed in standard list. It will be only ueed to generate random tasks for users
	//
	// required: true
	IsDynamic bool `json:"is_dynamic"`
	// Task content
	//
	// required: true
	Content ExpectedTaskContent `json:"content"     validate:"required"`
}

// ExpectedTaskContentDto
//
// Represents expected task content.
//
// swagger:model
type ExpectedTaskContent struct {
	// Type of content
	//
	// example: LESSON
	Kind TaskContentType `json:"kind"            validate:"required,oneof=LESSON SINGLE_SELECTION MULTI_SELECTION"`
	// Task description
	//
	// required: true
	Description string `json:"description"     validate:"required"`
	// Possible answers
	//
	// Required if kind is SINGLE_SELECTION or MULTI_SELECTION.
	//
	// required: false
	// min length: 2
	// max length: 64
	Options []ExpectedTaskOption `json:"options"         validate:"required_if=Kind SINGLE_SELECTION,required_if=Kind MULTI_SELECTION,omitempty,min=2,max=64,dive"`
	// Index of correct answer
	//
	// Required if kind is SINGLE_SELECTION.
	//
	// required: false
	// min: 0
	// max: 63
	CorrectOption *int `json:"correct_option"  validate:"required_if=Kind SINGLE_SELECTION,omitempty,min=0,max=63"`
	// Indexes of correct answers
	//
	// Required if kind is MULTI_SELECTION.
	//
	// required: false
	// min length: 1
	// max length: 64
	CorrectOptions []int `json:"correct_options" validate:"required_if=Kind MULTI_SELECTION,omitempty"`
	// Task hints
	//
	// required: false
	// max length: 128
	Hints []ExpectedTaskHint `json:"hints"           validate:"omitempty,max=128,dive"`
}

type ExpectedTaskOption struct {
	// Option content
	//
	// required: true
	Content string `json:"content" validate:"required"`
}

type ExpectedTaskHint struct {
	// Hint content
	//
	// required: true
	Content string `json:"content" validate:"required"`
}

func ExpectedTechnologiesSliceToDomain(
	expectedTechnologies []ExpectedTechnology,
) []model.ExpectedTechnology {
	technologies := make([]model.ExpectedTechnology, len(expectedTechnologies))
	for i, expectedTechnology := range expectedTechnologies {
		technologies[i] = expectedTechnologyToDomain(expectedTechnology, i)
	}
	return technologies
}

func expectedTechnologyToDomain(
	expectedTechnology ExpectedTechnology,
	position int,
) model.ExpectedTechnology {
	sections := make([]model.ExpectedSection, len(expectedTechnology.ExpectedSections))
	for i, expectedSection := range expectedTechnology.ExpectedSections {
		sections[i] = expectedSectionToDomain(expectedSection, i, expectedTechnology.Id)
	}
	return model.ExpectedTechnology{
		Technology: model.Technology{
			Id:          expectedTechnology.Id,
			Title:       expectedTechnology.Name,
			Description: expectedTechnology.Description,
			ImageUrl:    expectedTechnology.ImageUrl,
			Position:    position,
		},
		ExpectedSections: sections,
	}
}

func expectedSectionToDomain(
	expectedSection ExpectedSection,
	position int,
	parentTechnologyId string,
) model.ExpectedSection {
	tasks := make([]model.Task, len(expectedSection.Tasks))
	nextPosition := 0
	for i, expectedTask := range expectedSection.Tasks {
		var position *int
		if !expectedTask.IsDynamic {
			currentPosition := nextPosition
			position = &currentPosition
			nextPosition++
		}
		tasks[i] = expectedTaskToDomain(expectedTask, position, expectedSection.Id)
	}
	return model.ExpectedSection{
		Section: model.Section{
			Id:           expectedSection.Id,
			TechnologyId: parentTechnologyId,
			Title:        expectedSection.Name,
			Description:  expectedSection.Description,
			ImageUrl:     expectedSection.ImageUrl,
			Position:     position,
		},
		ExpectedTasks: tasks,
	}
}

func expectedTaskToDomain(
	expectedTask ExpectedTask,
	position *int,
	parentSectionId string,
) model.Task {
	content := contentToDomain(expectedTask.Content)

	return model.Task{
		Id:          expectedTask.Id,
		SectionId:   parentSectionId,
		Title:       expectedTask.Name,
		Description: expectedTask.Description,
		ImageUrl:    expectedTask.ImageUrl,
		Position:    position,
		Difficulty:  expectedTask.Difficulty,
		IsPublic:    expectedTask.IsPublic,
		Content:     &content,
	}
}

func contentToDomain(content ExpectedTaskContent) model.TaskContent {
	switch content.Kind {
	case TaskContentTypeLesson:
		return model.TaskContent{
			Lesson: &model.LessonTaskContent{
				Description: content.Description,
				Hints:       hintsToDomain(content.Hints),
			},
		}
	case TaskContentTypeSingleSelection:
		return model.TaskContent{
			SingleSelection: &model.SingleSelectionTaskContent{
				Description:     content.Description,
				Options:         optionsToDomain(content.Options),
				CorrectOptionId: *content.CorrectOption,
				Hints:           hintsToDomain(content.Hints),
			},
		}
	case TaskContentTypeMultiSelection:
		return model.TaskContent{
			MultiSelection: &model.MultiSelectionTaskContent{
				Description:      content.Description,
				Options:          optionsToDomain(content.Options),
				CorrectOptionIds: content.CorrectOptions,
				Hints:            hintsToDomain(content.Hints),
			},
		}
	default:
		return model.TaskContent{}
	}
}

func hintsToDomain(hints []ExpectedTaskHint) []model.Hint {
	result := make([]model.Hint, len(hints))
	for i, hint := range hints {
		result[i] = model.Hint{
			Id:      i,
			Content: hint.Content,
		}
	}
	return result
}

func optionsToDomain(options []ExpectedTaskOption) []model.Option {
	result := make([]model.Option, len(options))
	for i, option := range options {
		result[i] = model.Option{
			Id:      i,
			Content: option.Content,
		}
	}
	return result
}
