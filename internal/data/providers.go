package data

import (
	"context"

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
		NewDB,
		NewStoregeClient,
		NewBucketHandle,
		fx.Annotate(adapter.NewPgProfileRepository, fx.As(new(port.ProfileRepository))),
		fx.Annotate(adapter.NewBucketFileRepository, fx.As(new(port.FileRepository))),
	}
}

func NewDB(config *config.Config) (*db.Queries, error) {
	dbpool, err := pgxpool.New(context.Background(), config.PgUrl)
	if err != nil {
		return nil, err
	}
	return db.New(dbpool), nil
}

func NewStoregeClient() (*storage.Client, error) {
	return storage.NewClient(context.Background())
}

func NewBucketHandle(config *config.Config, client *storage.Client) *storage.BucketHandle {
	return client.Bucket(config.BucketName)
}
