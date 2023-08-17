package usecase

import (
	"context"
	"io"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type UploadProfileAvatar struct {
	profileRepository port.ProfileRepository
	fileRepository    port.FileRepository
}

func NewUploadProfileAvatar(
	profileRepository port.ProfileRepository,
	fileRepository port.FileRepository,
) *UploadProfileAvatar {
	return &UploadProfileAvatar{
		profileRepository: profileRepository,
		fileRepository:    fileRepository,
	}
}

func (u *UploadProfileAvatar) Invoke(ctx context.Context, imageReader io.Reader) error {
	profileId := util.ExtractCurrentUserId(ctx)
	url, err := u.fileRepository.UploadProfileAvatar(ctx, imageReader, *profileId)
	if err != nil {
		return err
	}
	err = u.profileRepository.UpdateProfileAvatar(ctx, *profileId, url)
	if err != nil {
		return err
	}

	return nil
}
