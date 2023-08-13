package dto

import (
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
)

// ErrorDto
//
// swagger:model
type Error struct {
	// Error code - for programmatic error handling
	//
	// example: internal_server_error
	// required: true
	Code string `json:"code"`
	// Error message - human readable
	//
	// example: Internal server error
	// required: true
	Message string `json:"message"`
	// Additional arguments
	//
	// example: {"arg1": "value1", "arg2": "value2"}
	// required: false
	Args map[string]interface{} `json:"args,omitempty"`

	StatusCode int `json:"-"`
}

func ErrorFromModel(err model.Error) Error {
	message := mapTypeToMessage(err)
	if message == "" {
		message = err.Error()
	}

	statusCode := mapTypeToStatusCode(err)
	if statusCode == 0 {
		statusCode = 500
	}

	return Error{
		Code:       string(err.Type),
		Message:    message,
		Args:       err.Args,
		StatusCode: statusCode,
	}
}

func mapTypeToMessage(err model.Error) string {
	switch err.Type {
	case usecase.ErrorTypeUnknown:
		return "Internal server error"
	case usecase.ErrorTypeUnauthorized:
		return "Unauthorized"
	case usecase.ErrorTypeNotFound:
		return "Not found"
	}
	return err.Error()
}

func mapTypeToStatusCode(err model.Error) int {
	switch err.Type {
	case usecase.ErrorTypeUnknown:
		return 500
	case usecase.ErrorTypeUnauthorized:
		return 401
	case usecase.ErrorTypeNotFound:
		return 404
	}
	return 500
}
