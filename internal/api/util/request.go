package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
)

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
	kind := reflect.TypeOf(dto).Kind()
	var err error
	switch kind {
	case reflect.Slice, reflect.Array:
		err = validate.Var(dto, "omitempty,dive")
	case reflect.Ptr, reflect.Struct:
		err = validate.Struct(dto)
	}

	if err == nil {
		return nil
	}

	slog.InfoContext(ctx, "Validation error", "err", err, "request_body", dto)

	var valError validator.ValidationErrors
	if errors.As(err, &valError) {
		args := make(map[string]interface{})

		for _, e := range valError {
			i := strings.Index(e.Namespace(), ".")
			field := e.Namespace()[i+1:]
			msg := fmt.Sprintf("%v %v", e.Tag(), e.Param())
			args[field] = strings.Trim(msg, " ")
		}

		return failure.NewWithArgs(failure.FailureTypeInvalidInput, args)
	}

	return failure.NewWithArgs(failure.FailureTypeInvalidInput, map[string]interface{}{
		"message": err.Error(),
	})
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
