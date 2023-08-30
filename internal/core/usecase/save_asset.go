package usecase

import (
	"context"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

type SaveAsset struct {
	contentRepository port.ContentRepository
	fileRepository    port.FileRepository
}

func NewSaveAsset(
	contentRepository port.ContentRepository,
	fileRepository port.FileRepository,
) *SaveAsset {
	return &SaveAsset{
		contentRepository: contentRepository,
		fileRepository:    fileRepository,
	}
}

func (s *SaveAsset) Invoke(ctx context.Context, assetReader io.Reader) (model.Asset, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to generate uuid", "err", err)

		return model.Asset{}, err
	}

	assetUrl, err := s.fileRepository.UploadContentAsset(ctx, assetReader, id.String())
	if err != nil {
		return model.Asset{}, err
	}

	asset, err := s.contentRepository.SaveAsset(ctx, id.String(), assetUrl)
	if err != nil {
		s.fileRepository.DeleteContentAsset(ctx, assetUrl)
		return model.Asset{}, err
	}

	return asset, nil
}
