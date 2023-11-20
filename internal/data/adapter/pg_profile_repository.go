package adapter

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/mapper"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

const (
	duplicateKeyErrorCode = "23505"
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

func (r *PgProfileRepository) UpsertProfile(
	ctx context.Context,
	id string,
	params model.CreateProfileParams,
) error {
	queryParams := mapper.CreateProfileParamsFromModel(id, params)
	_, err := r.queries.CreateProfile(ctx, queryParams)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == duplicateKeyErrorCode {
			return port.NameNotUniqueError
		}
		slog.ErrorContext(ctx, "Failed to upsert profile", "err", err)
		return err
	}

	return nil
}

func (r *PgProfileRepository) SetProfileAvatar(
	ctx context.Context,
	profileId string,
	url *string,
) error {
	err := r.queries.UpdateProfileAvatar(ctx, db.UpdateProfileAvatarParams{
		ID:        profileId,
		AvatarUrl: util.ToPgString(url),
	})
	if err != nil {
		slog.ErrorContext(ctx, "Failed to update profile avatar", "err", err)
		return err
	}

	return nil
}
