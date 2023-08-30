package mapper

import (
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/db"
	"github.com/LazyCodeTeam/just-code-backend/internal/data/util"
)

func AssetToDomain(asset db.Asset) model.Asset {
	return model.Asset{
		Id:  util.FromPgUUID(asset.ID),
		Url: asset.Url,
	}
}
