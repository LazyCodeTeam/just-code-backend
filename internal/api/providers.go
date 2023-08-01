package api

import (
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/handler"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(
			handler.NewHealthHandler,
			fx.ResultTags(`group:"routes"`),
		),
	}
}
