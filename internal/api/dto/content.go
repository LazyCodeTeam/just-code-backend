package dto

type TaskContentType string

const (
	TaskContentTypeLesson          TaskContentType = "LESSON"
	TaskContentTypeSingleSelection TaskContentType = "SINGLE_SELECTION"
	TaskContentTypeMultiSelection  TaskContentType = "MULTI_SELECTION"
)
