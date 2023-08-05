package model

type AuthType string

const (
	AuthTypeAnonymous AuthType = "anonymous"
	AuthTypeEmail     AuthType = "email"
)

type AuthData struct {
	Id       string
	Type     AuthType
	Email    *string
	Verified bool
}

func IsAnonymous(authData AuthData) bool {
	return authData.Type == AuthTypeAnonymous
}
