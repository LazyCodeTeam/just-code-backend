package data

import (
	"context"
	"log/slog"

	"cloud.google.com/go/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	"github.com/LazyCodeTeam/just-code-backend/internal/config"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/adapter"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
)

func Providers() []interface{} {
	return []interface{}{
		newConnectionPool,
		newDB,
		newStoregeCliente,
		newBucketHandle,
		fx.Annotate(adapter.NewPgProfileRepository, fx.As(new(port.ProfileRepository))),
		fx.Annotate(adapter.NewBucketFileRepository, fx.As(new(port.FileRepository))),
		fx.Annotate(adapter.NewPgContentRepository, fx.As(new(port.ContentRepository))),
		fx.Annotate(adapter.NewPgTransactionFactory, fx.As(new(port.TransactionFactory))),
	}
}

func newConnectionPool(config *config.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), config.PgUrl)
	if err != nil {
		slog.Error("Failed to connect to database", "err", err)
		return nil, err
	}

	return dbpool, nil
}

func newDB(config *config.Config, pool *pgxpool.Pool) (*db.Queries, error) {
	return db.New(pool), nil
}

func newStoregeCliente() (*storage.Client, error) {
	return storage.NewClient(context.Background())
}

func newBucketHandle(config *config.Config, client *storage.Client) *storage.BucketHandle {
	return client.Bucket(config.BucketName)
}
