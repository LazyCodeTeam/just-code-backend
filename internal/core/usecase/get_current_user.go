package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type GetCurrentUser struct {
	profileRepository port.ProfileRepository
}

func NewGetCurrentUser(profileRepository port.ProfileRepository) *GetCurrentUser {
	return &GetCurrentUser{profileRepository: profileRepository}
}

func (u *GetCurrentUser) Invoke(ctx context.Context) (model.Profile, error) {
	id := util.ExtractCurrentUserId(ctx)

	profile, err := u.profileRepository.GetProfileById(ctx, *id)
	if err != nil {
		return model.Profile{}, model.NewError(ErrorTypeUnknown)
	}

	if profile == nil {
		return model.Profile{}, model.NewError(ErrorTypeNotFound)
	}

	return *profile, nil
}
