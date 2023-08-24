package model

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
	Content     TaskContent
}

type TaskContent struct{}
