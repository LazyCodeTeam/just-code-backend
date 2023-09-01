package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type GetTasks struct {
	contentRepository port.ContentRepository
}

func NewGetTasks(contentRepository port.ContentRepository) *GetTasks {
	return &GetTasks{contentRepository: contentRepository}
}

func (u *GetTasks) Invoke(
	ctx context.Context,
	sectionId string,
) ([]model.Task, error) {
	authData := util.ExtractAuthData(ctx)
	_ = authData

	tasks, err := u.contentRepository.GetSectionTasks(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	if authData.Type == model.AuthTypeAnonymous {
		for i := range tasks {
			if !tasks[i].IsPublic {
				tasks[i].Content = nil
			}
		}

		return tasks, nil
	}

	// Handle historical answers
	return tasks, nil
}
