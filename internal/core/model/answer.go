package model

type Answer struct {
	TaskId     string
	AnswerData AnswerData
}

type AnswerData struct {
	SingleAnswer *SingeAnswerData
	MultiAnswer  *MultiAnswerData
}

type SingeAnswerData struct {
	Answer int
}

type MultiAnswerData struct {
	Answers []int
}

type AnswerResult string

const (
	AnswerResultFirstValid AnswerResult = "FIRST_VALID"
	AnswerResultValid      AnswerResult = "VALID"
	AnswerResultInvalid    AnswerResult = "INVALID"
)

type HistoricalAnswer struct {
	Id           string
	ProfileId    string
	TaskId       string
	AnswerResult AnswerResult
	created_at   string
}
