package api

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/middleware"
)

func Providers() []interface{} {
	return []interface{}{
		middleware.NewAuthTokenValidator,
		handler.NewHealthHandler,
		createValidator,
		fx.Annotate(
			handler.NewProfileHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(handler.Handler)),
		),
		fx.Annotate(
			handler.NewContentHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(handler.Handler)),
		),
		fx.Annotate(
			handler.NewAnswerHandler,
			fx.ResultTags(`group:"routes"`),
			fx.As(new(handler.Handler)),
		),
		handler.NewAdminContentHandler,
	}
}

func createValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate
}
