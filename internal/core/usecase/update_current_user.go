package usecase

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/port"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type UpdateCurrentProfile struct {
	profileRepository port.ProfileRepository
}

func NewUpdateCurrentProfile(profileRepository port.ProfileRepository) *UpdateCurrentProfile {
	return &UpdateCurrentProfile{profileRepository: profileRepository}
}

func (u *UpdateCurrentProfile) Invoke(ctx context.Context, params model.CreateProfileParams) error {
	id := util.ExtractCurrentUserId(ctx)

	return u.profileRepository.UpsertProfile(ctx, *id, params)
}
