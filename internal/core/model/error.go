package model

import "fmt"

type ErrorType string

type Error struct {
	Type ErrorType
	Args map[string]interface{}
}

func NewError(t ErrorType) *Error {
	return &Error{
		Type: t,
	}
}

func NewErrorWithArgs(t ErrorType, a map[string]interface{}) *Error {
	return &Error{
		Type: t,
		Args: a,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v - %#v", e.Type, e.Args)
}
