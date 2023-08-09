package model

import "time"

type Profile struct {
	Nick      string
	FirstName *string
	LastName  *string
	AvatarUrl *string
	UpdatedAt time.Time
	CreatedAt time.Time
}
