package adapter

import "github.com/LazyCodeTeam/just-code-backend/internal/data/db"

type PgContentRepository struct {
	queries *db.Queries
}

func NewPgContentRepository(queries *db.Queries) *PgContentRepository {
	return &PgContentRepository{queries: queries}
}
