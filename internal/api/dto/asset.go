package dto

import "github.com/LazyCodeTeam/just-code-backend/internal/core/model"

// AssetDto
//
// # Represents asset
//
// swagger:model
type Asset struct {
	// Asset id
	//
	// required: true
	// format: uuid
	Id string `json:"id"`
	// Asset url
	//
	// required: true
	Url string `json:"url"`
}

func AssetFromDomain(asset model.Asset) Asset {
	return Asset{
		Id:  asset.Id,
		Url: asset.Url,
	}
}
