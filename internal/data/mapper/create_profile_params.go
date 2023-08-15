package mapper

import (
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

func CreateProfileParamsFromModel(
	id string,
	params model.CreateProfileParams,
) db.CreateProfileParams {
	return db.CreateProfileParams{
		ID:        id,
		Name:      params.Nick,
		FirstName: util.ToDbString(params.FirstName),
		LastName:  util.ToDbString(params.LastName),
		AvatarUrl: util.ToDbString(params.AvatarUrl),
	}
}
