package dto

import "github.com/LazyCodeTeam/just-code-backend/internal/core/model"

// AnswerDto
//
// Represents answer to the task.
//
// swagger:model
type Answer struct {
	// Task id
	//
	// required: true
	// format: uuid
	TaskId string `json:"task_id"       validate:"required"`
	// Kind of answer
	//
	// required: true
	// enum: EMPTY, SINGLE, MULTI
	Kind string `json:"kind"          validate:"required,oneof=EMPTY SINGLE MULTI"`
	// Answer - required if kind is SINGLE
	//
	// required: false
	SingleAnswer *int `json:"single_answer" validate:"required_if=Kind SINGLE,omitempty,gte=0"`
	// Answers - required if kind is MULTI
	//
	// required: false
	MultiAnswer *[]int `json:"multi_answer"  validate:"required_if=Kind MULTI,omitempty,dive,gte=0"`
}

// AnswerResultDto
//
// Represents result of answer validation.
//
// swagger:model
type AnswerResult struct {
	// Result of answer validation
	//
	// required: true
	// enum: FIRST_VALID, VALID, INVALID
	Result string `json:"result"`
}

func AnswerResultFromModel(result model.AnswerResult) AnswerResult {
	return AnswerResult{Result: string(result)}
}

func AnswerToModel(answer Answer) model.Answer {
	var data model.AnswerData
	if answer.Kind == "SINGLE" {
		data = model.AnswerData{SingleAnswer: &model.SingeAnswerData{Answer: *answer.SingleAnswer}}
	}
	if answer.Kind == "MULTI" {
		data = model.AnswerData{MultiAnswer: &model.MultiAnswerData{Answers: *answer.MultiAnswer}}
	}
	return model.Answer{
		TaskId:     answer.TaskId,
		AnswerData: data,
	}
}
