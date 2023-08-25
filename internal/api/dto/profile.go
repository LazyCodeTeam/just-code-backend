package dto

import (
	"time"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

// ProfileDto
//
// Represents user profile.
//
// swagger:model
type Profile struct {
	// User id
	//
	// example: 123e4567-e89b-12d3-a456-426614174000
	// required: true
	Id string `json:"id"`
	// User nickname
	//
	// example: johndoe
	// required: true
	Nickname string `json:"nickname"`
	// User first name
	//
	// example: John
	// required: false
	FirstName *string `json:"first_name,omitempty"`
	// User last name
	//
	// example: Doe
	// required: false
	LastName *string `json:"last_name,omitempty"`
	// User avatar url
	//
	// example: https://example.com/avatar.png
	// required: false
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// User created at
	//
	// example: 2021-01-01T00:00:00Z
	// required: true
	CreatedAt time.Time `json:"created_at"`
}

func ProfileFromDomain(model model.Profile) Profile {
	return Profile{
		Id:        model.Id,
		Nickname:  model.Nick,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		AvatarUrl: model.AvatarUrl,
		CreatedAt: model.CreatedAt,
	}
}
