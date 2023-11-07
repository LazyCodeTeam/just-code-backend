package usecase

import (
	"context"
	"fmt"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
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

	err := u.profileRepository.UpsertProfile(ctx, *id, params)
	if err == port.NameNotUniqueError {
		return failure.NewStateFailure(failure.FailureTypeUsernameNotUnique, err)
	}
	if err != nil {
		return fmt.Errorf("Failed to update profile: %w", err)
	}

	return nil
}
