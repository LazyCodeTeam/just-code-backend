package dto

import (
	"net/http"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

// Error
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
	code := typeToCode[err.Type]
	if code == "" {
		code = "internal_server_error"
	}

	message := typeToMessage[err.Type]
	if message == "" {
		message = err.OriginalError.Error()
	}

	statusCode := typeToStatusCode[err.Type]
	if statusCode == 0 {
		statusCode = 500
	}

	return Error{
		Code:       code,
		Message:    message,
		Args:       err.Args,
		StatusCode: statusCode,
	}
}

var typeToCode = map[string]string{
	model.ErrorUnknown:      "internal_server_error",
	model.ErrorUnauthorized: "unauthorized",
}

var typeToMessage = map[string]string{
	model.ErrorUnknown:      "Internal server error",
	model.ErrorUnauthorized: "Unauthorized",
}

var typeToStatusCode = map[string]int{
	model.ErrorUnknown:      http.StatusInternalServerError,
	model.ErrorUnauthorized: http.StatusUnauthorized,
}
