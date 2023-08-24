package dto

import "github.com/LazyCodeTeam/just-code-backend/internal/core/model"

// CreateProfileParamsDto
//
// swagger:model
type CreateProfileParams struct {
	// User nickname
	//
	// example: johndoe
	// required: true
	Nickname string `json:"nickname"   validate:"required,min=3,max=48"`
	// User first name
	//
	// example: John
	// required: false
	FirstName *string `json:"first_name" validate:"omitempty,min=2,max=64"`
	// User last name
	//
	// example: Doe
	// required: false
	LastName *string `json:"last_name"  validate:"omitempty,min=2,max=64"`
}

func CreateProfileParamsToDomain(dto CreateProfileParams) model.CreateProfileParams {
	return model.CreateProfileParams{
		Nick:      dto.Nickname,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}
