package api

import (
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
	"github.com/LazyCodeTeam/just-code-backend/internal/api/middleware"
)

func Providers() []interface{} {
	return []interface{}{
		middleware.NewAuthTokenValidator,
		handler.NewHealthHandler,
		fx.Annotate(
			handler.NewCurrentProfileHandler,
			fx.ResultTags(`group:"routes"`),
		),
	}
}
