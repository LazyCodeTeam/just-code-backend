package model

import "fmt"

const (
	ErrorUnknown      = "unknown_error"
	ErrorUnauthorized = "unauthorized_error"
)

type Error struct {
	OriginalError error
	Type          string
	Args          map[string]interface{}
}

func (e *Error) Error() string {
	if e.OriginalError == nil {
		return fmt.Sprintf("%s", e.Type)
	}

	return fmt.Sprintf("%s: %s", e.Type, e.OriginalError.Error())
}

func NewUnknownError(original error) *Error {
	return &Error{
		OriginalError: original,
		Type:          ErrorUnknown,
	}
}

func NewUnauthorizedError(original error) *Error {
	return &Error{
		OriginalError: original,
		Type:          ErrorUnauthorized,
	}
}
