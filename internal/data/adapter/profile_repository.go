package adapter

import (
	"context"

	"github.com/jackc/pgx/v5"
	"log/slog"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/mapper"
)

type PgProfileRepository struct {
	queries *db.Queries
}

func NewPgProfileRepository(queries *db.Queries) *PgProfileRepository {
	return &PgProfileRepository{queries: queries}
}

func (r *PgProfileRepository) GetProfileById(
	ctx context.Context,
	id string,
) (*model.Profile, error) {
	profile, err := r.queries.GetProfileById(ctx, id)
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		slog.ErrorContext(ctx, "Failed to get profile by id", "err", err)
		return nil, err
	}

	domainProfile := mapper.ProfleToDomain(profile)
	return &domainProfile, nil
}
