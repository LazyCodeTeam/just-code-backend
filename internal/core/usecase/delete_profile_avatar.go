package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type DeleteProfileAvatar struct {
	profileRepository port.ProfileRepository
	fileRepository    port.FileRepository
}

func NewDeleteProfileAvatar(
	profileRepository port.ProfileRepository,
	fileRepository port.FileRepository,
) *DeleteProfileAvatar {
	return &DeleteProfileAvatar{
		profileRepository: profileRepository,
		fileRepository:    fileRepository,
	}
}

func (u *DeleteProfileAvatar) Invoke(ctx context.Context) error {
	profileId := util.ExtractCurrentUserId(ctx)
	err := u.fileRepository.DeleteProfileAvatar(ctx, *profileId)
	if err != nil {
		return err
	}

	return u.profileRepository.SetProfileAvatar(ctx, *profileId, nil)
}
