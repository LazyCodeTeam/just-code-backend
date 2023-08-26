package util

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
)

func DeserializeAndValidateBodySlice[T any](
	r *http.Request,
	validate *validator.Validate,
) ([]T, error) {
	dtos, err := DeserializeBody[[]T](r)
	if err != nil {
		return dtos, err
	}
	for _, item := range dtos {
		err = Validate[T](r.Context(), item, validate)
		if err != nil {
			return dtos, err
		}
	}
	return dtos, nil
}

func DeserializeAndValidateBody[T any](r *http.Request, validate *validator.Validate) (T, error) {
	dto, err := DeserializeBody[T](r)
	if err != nil {
		return dto, err
	}

	err = Validate[T](r.Context(), dto, validate)

	if err != nil {
		return dto, err
	}
	return dto, nil
}

func Validate[T any](ctx context.Context, dto T, validate *validator.Validate) error {
	err := validate.Struct(dto)
	if err != nil {
		slog.WarnContext(ctx, "Validation error", "err", err, "request_body", dto)
		return failure.NewWithArgs(failure.FailureTypeInvalidInput, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return nil
}

func DeserializeBody[T any](r *http.Request) (T, error) {
	var dto T
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		slog.WarnContext(r.Context(), "Deserialization error", "err", err)
		return dto, failure.NewWithArgs(failure.FailureTypeInvalidInput, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return dto, nil
}
