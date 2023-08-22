package port

import "context"

type TransactionFactory interface {
	Begin(ctx context.Context) (Transaction, error)
}

type Transaction interface {
	ContentRepository(ctx context.Context) ContentRepository

	Commit(ctx context.Context) error

	Rollback(ctx context.Context) error
}
