package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/config"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/adapter"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
)

func Providers() []interface{} {
	return []interface{}{
		NewDB,
		fx.Annotate(adapter.NewPgProfileRepository, fx.As(new(port.ProfileRepository))),
	}
}

func NewDB(config *config.Config) (*db.Queries, error) {
	dbpool, err := pgxpool.New(context.Background(), config.PgUrl)
	if err != nil {
		return nil, err
	}
	return db.New(dbpool), nil
}
