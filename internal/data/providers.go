package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LazyCodeTeam/just-code-backend/internal/config"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
)

func Providers() []interface{} {
	return []interface{}{
		NewDB,
	}
}

func NewDB(config *config.Config) (*db.Queries, error) {
	dbpool, err := pgxpool.New(context.Background(), config.PgUrl)
	if err != nil {
		return nil, err
	}
	return db.New(dbpool), nil
}
