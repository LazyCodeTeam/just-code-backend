package adapter

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
)

type PgTransactionFactory struct {
	db *pgxpool.Pool
}

func NewPgTransactionFactory(db *pgxpool.Pool) *PgTransactionFactory {
	return &PgTransactionFactory{db: db}
}

func (f *PgTransactionFactory) Begin(ctx context.Context) (port.Transaction, error) {
	tx, err := f.db.Begin(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to begin transaction: %v", "err", err)

		return nil, err
	}
	queries := db.New(tx)

	return &PgTransaction{tx: tx, queries: queries}, nil
}

type PgTransaction struct {
	tx       pgx.Tx
	queries  *db.Queries
	finished bool
}

func (t *PgTransaction) ContentRepository(ctx context.Context) port.ContentRepository {
	return NewPgContentRepository(t.queries)
}

func (t *PgTransaction) AnswerRepository(ctx context.Context) port.AnswerRepository {
	return NewPgAnswerRepository(t.queries)
}

func (t *PgTransaction) Commit(ctx context.Context) error {
	if t.finished {
		slog.WarnContext(ctx, "Failed to commit transaction: transaction already finished")
		return nil
	}
	defer func() {
		t.finished = true
	}()

	err := t.tx.Commit(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to commit transaction", "err", err)
		return err
	}

	return nil
}

func (t *PgTransaction) Rollback(ctx context.Context) {
	if t.finished {
		slog.DebugContext(ctx, "Skipping rollback: transaction already commited")
		return
	}
	defer func() {
		t.finished = true
	}()

	err := t.tx.Rollback(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to rollback transaction: %v", "err", err)
		return
	}
}
