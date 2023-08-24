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
		FirstName: util.FromPgString(profile.FirstName),
		LastName:  util.FromPgString(profile.LastName),
		AvatarUrl: util.FromPgString(profile.AvatarUrl),
		UpdatedAt: profile.UpdatedAt.Time,
		CreatedAt: profile.CreatedAt.Time,
	}
}
