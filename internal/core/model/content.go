package model

import "time"

type Technology struct {
	Id          string
	Title       string
	Description *string
	ImageUrl    *string
	Position    int
}

type Section struct {
	Id           string
	TechnologyId string
	Title        string
	Description  *string
	ImageUrl     *string
	Position     int
}

type Task struct {
	Id          string
	SectionId   string
	Title       string
	Description *string
	ImageUrl    *string
	Position    *int
	Difficulty  int
	IsPublic    bool
	Content     *TaskContent
	DoneAt      *time.Time
}

type TaskContent struct {
	Lesson          *LessonTaskContent
	SingleSelection *SingleSelectionTaskContent
	MultiSelection  *MultiSelectionTaskContent
}

type LessonTaskContent struct {
	Description string
	Hints       []Hint
}

type SingleSelectionTaskContent struct {
	Description     string
	Options         []Option
	CorrectOptionId int
	Hints           []Hint
}

type MultiSelectionTaskContent struct {
	Description      string
	Options          []Option
	CorrectOptionIds []int
	Hints            []Hint
}

type Hint struct {
	Id      int
	Content string
}

type Option struct {
	Id      int
	Content string
}

type SectionPreview struct {
	Id    string
	Title string
}

type TaskPreview struct {
	Id       string
	Title    string
	IsPublic bool
	DoneAt   *time.Time
}

type TechnologyWithSectionsPreview struct {
	Technology
	Sections []SectionPreview
}

type SectionWithTasksPreview struct {
	Section
	Tasks []TaskPreview
}
