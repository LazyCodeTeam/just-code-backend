package dto

import "github.com/LazyCodeTeam/just-code-backend/internal/core/model"

// ExpectedTechnologyDto - represents desired state of technology.
// Will be compared with actual state and required changes will be applied
//
// swagger:model
type ExpectedTechnology struct {
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
	// Expended sections
	//
	// required: true
	// min length: 1
	ExpectedSections []ExpectedSection `json:"sections"`
}

// ExpectedSectionDto - represents desired state of section.
// Will be compared with actual state and required changes will be applied
//
// swagger:model
type ExpectedSection struct {
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
	// Section description
	//
	// min length: 1
	// required: false
	Description *string `json:"description"`
	// Section image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url"`
	// Expended tasks
	//
	// required: true
	// min length: 1
	ExpectedTasks []ExpectedTask `json:"tasks"`
}

// ExpectedTaskDto - represents desired state of task in section.
// Will be compared with actual state and required changes will be applied
//
// swagger:model
type ExpectedTask struct {
	// Task id -- UUID
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Task title
	//
	// required: true
	// min length: 1
	// max length: 1024
	Title string `json:"title"`
	// Task description
	//
	// min length: 1
	// required: false
	Description *string `json:"description"`
	// Task image url
	//
	// required: false
	// min length: 1
	ImageUrl *string `json:"image_url"`
	// Task difficulty
	//
	// required: true
	// min: 1
	// max: 10
	Difficulty int `json:"difficulty"`
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
	Content TaskContent `json:"content"`
}

// TaskContentDto
//
// swagger:model
type TaskContent struct{}

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
		sections[i] = expectedSectionToDomain(expectedSection, i)
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

func expectedSectionToDomain(expectedSection ExpectedSection, position int) model.ExpectedSection {
	tasks := make([]model.Task, len(expectedSection.ExpectedTasks))
	i := 0
	for _, expectedTask := range expectedSection.ExpectedTasks {
		currentIndex := i
		var position *int
		if !expectedTask.IsDynamic {
			position = &currentIndex
			i++
		}
		tasks[i] = expectedTaskToDomain(expectedTask, position)
	}
	return model.ExpectedSection{
		Section: model.Section{
			Id:          expectedSection.Id,
			Title:       expectedSection.Name,
			Description: expectedSection.Description,
			ImageUrl:    expectedSection.ImageUrl,
			Position:    position,
		},
		ExpectedTasks: tasks,
	}
}

func expectedTaskToDomain(expectedTask ExpectedTask, position *int) model.Task {
	return model.Task{
		Id:          expectedTask.Id,
		Title:       expectedTask.Title,
		Description: expectedTask.Description,
		ImageUrl:    expectedTask.ImageUrl,
		Position:    position,
		Difficulty:  expectedTask.Difficulty,
		IsPublic:    expectedTask.IsPublic,
		Content:     contentToDomain(expectedTask.Content),
	}
}

func contentToDomain(content TaskContent) model.TaskContent {
	return model.TaskContent{}
}
