package util

import (
	"encoding/json"
	"net/http"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/dto"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

func WriteError(writer http.ResponseWriter, err error) {
	var errorModel *model.Error
	e, ok := err.(*model.Error)
	if ok {
		errorModel = e
	} else {
		errorModel = model.NewUnknownError(err)
	}

	dto := dto.ErrorFromModel(*errorModel)
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
		json.NewEncoder(writer).Encode(response)
	}
}
