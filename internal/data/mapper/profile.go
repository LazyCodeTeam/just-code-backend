package mapper

import (
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

func ProfleToDomain(profile db.Profile) model.Profile {
	return model.Profile{
		Id:        profile.ID,
		Nick:      profile.Name,
		FirstName: util.UnwrapDbString(profile.FirstName),
		LastName:  util.UnwrapDbString(profile.LastName),
		AvatarUrl: util.UnwrapDbString(profile.AvatarUrl),
		UpdatedAt: profile.UpdatedAt.Time,
		CreatedAt: profile.CreatedAt.Time,
	}
}
