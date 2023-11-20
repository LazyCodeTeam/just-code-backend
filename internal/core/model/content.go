package model

import (
	"time"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
)

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
	Title       *string
	Description *string
	ImageUrl    *string
	Position    *int
	Difficulty  int
	IsPublic    bool
	Content     *TaskContent
	DoneAt      *time.Time
}

func (t *Task) IsAnswerValid(answer Answer) (AnswerResult, error) {
	var isValid bool
	var err error
	switch {
	case t.Content.Lesson != nil:
		isValid, err = t.Content.Lesson.IsAnswerValid(answer)
	case t.Content.SingleSelection != nil:
		isValid, err = t.Content.SingleSelection.IsAnswerValid(answer)
	case t.Content.MultiSelection != nil:
		isValid, err = t.Content.MultiSelection.IsAnswerValid(answer)
	case t.Content.LinesArrangement != nil:
		isValid, err = t.Content.LinesArrangement.IsAnswerValid(answer)
	}
	if err != nil {
		return AnswerResultInvalid, err
	}

	if isValid && t.DoneAt == nil {
		return AnswerResultFirstValid, nil
	} else if isValid {
		return AnswerResultValid, nil
	} else {
		return AnswerResultInvalid, nil
	}
}

type TaskContent struct {
	Lesson           *LessonTaskContent
	SingleSelection  *SingleSelectionTaskContent
	MultiSelection   *MultiSelectionTaskContent
	LinesArrangement *LinesArrangementTaskContent
}

type LinesArrangementTaskContent struct {
	Description  string
	Lines        []Option
	CorrectOrder []int
	Hints        []Hint
}

func (l *LinesArrangementTaskContent) IsAnswerValid(answer Answer) (bool, error) {
	if answer.AnswerData.MultiAnswer == nil {
		return false, failure.NewInputFailure(
			failure.FailureTypeInvalidAnswerType,
			"expected_type",
			"multi_answer",
		)
	}

	if len(l.CorrectOrder) != len(answer.AnswerData.MultiAnswer.Answers) {
		return false, nil
	}

	for i, correctLineId := range l.CorrectOrder {
		if answer.AnswerData.MultiAnswer.Answers[i] != correctLineId {
			return false, nil
		}
	}

	return true, nil
}

type LessonTaskContent struct {
	Description string
	Hints       []Hint
}

func (l *LessonTaskContent) IsAnswerValid(answer Answer) (bool, error) {
	return true, nil
}

type SingleSelectionTaskContent struct {
	Description     string
	Options         []Option
	CorrectOptionId int
	Hints           []Hint
}

func (s *SingleSelectionTaskContent) IsAnswerValid(answer Answer) (bool, error) {
	if answer.AnswerData.SingleAnswer == nil {
		return false, failure.NewInputFailure(
			failure.FailureTypeInvalidAnswerType,
			"expected_type",
			"single_answer",
		)
	}

	return s.CorrectOptionId == answer.AnswerData.SingleAnswer.Answer, nil
}

type MultiSelectionTaskContent struct {
	Description      string
	Options          []Option
	CorrectOptionIds []int
	Hints            []Hint
}

func (m *MultiSelectionTaskContent) IsAnswerValid(answer Answer) (bool, error) {
	if answer.AnswerData.MultiAnswer == nil {
		return false, failure.NewInputFailure(
			failure.FailureTypeInvalidAnswerType,
			"expected_type",
			"multi_answer",
		)
	}

	if len(m.CorrectOptionIds) != len(answer.AnswerData.MultiAnswer.Answers) {
		return false, nil
	}

	correctOptionsMap := make(map[int]bool)
	for _, correctOptionId := range m.CorrectOptionIds {
		correctOptionsMap[correctOptionId] = true
	}
	for _, answerOptionId := range answer.AnswerData.MultiAnswer.Answers {
		if !correctOptionsMap[answerOptionId] {
			return false, nil
		}
		delete(correctOptionsMap, answerOptionId)
	}

	return len(correctOptionsMap) == 0, nil
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
	Title    *string
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
