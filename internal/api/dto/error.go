package dto

import (
	"net/http"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
)

// ErrorDto
//
// Represents error.
//
// swagger:model
type Error struct {
	// Error code - for programmatic error handling
	//
	// example: internal_server_error
	// required: true
	Code string `json:"code"`
	// Additional arguments
	//
	// example: {"arg1": "value1", "arg2": "value2"}
	// required: false
	Args map[string]interface{} `json:"args,omitempty"`

	StatusCode int `json:"-"`
}

func ErrorFromDomain(err failure.Failure) Error {
	statusCode := mapTypeToStatusCode(err)
	if statusCode == 0 {
		statusCode = 500
	}

	return Error{
		Code:       string(err.Type),
		Args:       err.Args,
		StatusCode: statusCode,
	}
}

func mapTypeToStatusCode(err failure.Failure) int {
	switch err.Type {
	case failure.FailureTypeUnknown:
		return http.StatusInternalServerError
	case failure.FailureTypeInvalidInput:
		return http.StatusBadRequest
	case failure.FailureTypeUnauthorized:
		return http.StatusUnauthorized
	case failure.FailureTypeNotFound:
		return http.StatusNotFound
	case failure.FailureTypeValueNotUnique:
		return http.StatusConflict
	case failure.FailureTypeUnsupportedFileType:
		return http.StatusUnsupportedMediaType
	}
	return http.StatusInternalServerError
}
