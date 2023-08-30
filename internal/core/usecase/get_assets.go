package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

type GetAssets struct {
	contentRepository port.ContentRepository
}

func NewGetAssets(
	contentRepository port.ContentRepository,
) *GetAssets {
	return &GetAssets{
		contentRepository: contentRepository,
	}
}

func (s *GetAssets) Invoke(ctx context.Context) ([]model.Asset, error) {
	return s.contentRepository.GetAssets(ctx)
}
