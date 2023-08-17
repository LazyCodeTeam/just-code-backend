package port

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

type ProfileRepository interface {
	GetProfileById(context.Context, string) (*model.Profile, error)

	UpdateProfileAvatar(ctx context.Context, profileId string, url string) error

	UpsertProfile(context.Context, string, model.CreateProfileParams) error
}
