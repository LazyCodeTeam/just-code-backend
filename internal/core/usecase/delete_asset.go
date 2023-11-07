package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

type DeleteAsset struct {
	contentRepository port.ContentRepository
	fileRepository    port.FileRepository
}

func NewDeleteAsset(
	contentRepository port.ContentRepository,
	fileRepository port.FileRepository,
) *DeleteAsset {
	return &DeleteAsset{
		contentRepository: contentRepository,
		fileRepository:    fileRepository,
	}
}

func (s *DeleteAsset) Invoke(ctx context.Context, id string) error {
	err := s.contentRepository.DeleteAsset(ctx, id)
	if err != nil {
		return err
	}

	err = s.fileRepository.DeleteContentAsset(ctx, id)
	if err == port.FileNotFoundError {
		return failure.NewNotFoundFailure(failure.FailureTypeFileNotFound, err, "id", id)
	}
	if err != nil {
		return err
	}

	return nil
}
