package util

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
)

func WriteError(writer http.ResponseWriter, err error) {
	var errorModel *failure.Failure
	e, ok := err.(*failure.Failure)
	if ok {
		errorModel = e
	} else {
		errorModel = failure.NewUnknownFailure(failure.FailureTypeUnknown, err)
	}

	dto := dto.ErrorFromDomain(*errorModel)
	WriteResponseJson(writer, dto, dto.StatusCode)
}

func WriteResponseJson(writer http.ResponseWriter, response interface{}, statusCode ...int) {
	writer.Header().Set("Content-Type", "application/json")
	if len(statusCode) > 0 {
		writer.WriteHeader(statusCode[0])
	} else {
		writer.WriteHeader(http.StatusOK)
	}
	if response != nil {
		err := json.NewEncoder(writer).Encode(response)
		if err != nil {
			slog.Error("Error while encoding response", "err", err)
		}
	}
}
