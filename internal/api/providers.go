package api

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/middleware"
)

func Providers() []interface{} {
	return []interface{}{
		middleware.NewAuthTokenValidator,
		handler.NewHealthHandler,
		validator.New,
		fx.Annotate(
			handler.NewProfileGetCurrentHandler,
			fx.ResultTags(`group:"routes"`),
		),
		fx.Annotate(
			handler.NewProfilePutCurrentHandler,
			fx.ResultTags(`group:"routes"`),
		),
	}
}
