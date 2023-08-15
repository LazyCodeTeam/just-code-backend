package model

import "time"

type Profile struct {
	Id        string
	Nick      string
	FirstName *string
	LastName  *string
	AvatarUrl *string
	UpdatedAt time.Time
	CreatedAt time.Time
}
