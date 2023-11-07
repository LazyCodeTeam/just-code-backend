package port

import (
	"context"
	"errors"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

var NameNotUniqueError = errors.New("name is not unique")

type ProfileRepository interface {
	GetProfileById(context.Context, string) (*model.Profile, error)

	SetProfileAvatar(ctx context.Context, profileId string, url *string) error

	UpsertProfile(context.Context, string, model.CreateProfileParams) error
}
