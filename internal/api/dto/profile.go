package dto

import "github.com/LazyCodeTeam/just-code-backend/internal/core/model"

// ProfileDto
//
// swagger:model
type Profile struct {
	// User nickname
	//
	// example: johndoe
	// required: true
	Nickname string `json:"nickname"`
	// User first name
	//
	// example: John
	// required: false
	FirstName *string `json:"first_name"`
	// User last name
	//
	// example: Doe
	// required: false
	LastName *string `json:"last_name"`
	// User avatar url
	//
	// example: https://example.com/avatar.png
	// required: false
	AvatarUrl *string `json:"avatar_url"`
}

func ProfileFromModel(model model.Profile) Profile {
	return Profile{
		Nickname:  model.Nick,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		AvatarUrl: model.AvatarUrl,
	}
}
