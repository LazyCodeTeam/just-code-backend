package adapter

import (
	"context"
	"log/slog"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/mapper"
)

type PgAnswerRepository struct {
	queries *db.Queries
}

func NewPgAnswerRepository(queries *db.Queries) *PgAnswerRepository {
	return &PgAnswerRepository{queries: queries}
}

func (r *PgAnswerRepository) SaveHistoricalAnswer(
	ctx context.Context,
	answer model.HistoricalAnswer,
) error {
	err := r.queries.InsertAnswer(ctx, mapper.InsertAnswerParamsFromModel(answer))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to insert answer", "err", err)
		return err
	}

	return nil
}
