package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
)

type GetSections struct {
	contentRepository port.ContentRepository
}

func NewGetSections(contentRepository port.ContentRepository) *GetSections {
	return &GetSections{contentRepository: contentRepository}
}

func (u *GetSections) Invoke(
	ctx context.Context,
	technologyId string,
) ([]model.SectionWithTasksPreview, error) {
	return u.contentRepository.GetSectionsWithTasksPreview(ctx, technologyId)
}
