package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

type GetTechnologies struct {
	contentRepository port.ContentRepository
}

func NewGetTechnologies(contentRepository port.ContentRepository) *GetTechnologies {
	return &GetTechnologies{contentRepository: contentRepository}
}

func (g *GetTechnologies) Invoke(
	ctx context.Context,
) ([]model.TechnologyWithSectionsPreview, error) {
	return g.contentRepository.GetTechnologiesWithSectionsPreview(ctx)
}
