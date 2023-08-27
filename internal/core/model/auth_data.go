package model

type AuthType string

const (
	AuthTypeAnonymous AuthType = "anonymous"
	AuthTypeEmail     AuthType = "email"
)

const (
	AuthRoleAdmin string = "admin"
)

type AuthData struct {
	Id       string
	Type     AuthType
	Email    *string
	Verified bool
	Roles    map[string]bool
}

func (a *AuthData) CanAccess(roles ...string) bool {
	if len(roles) == 0 {
		return true
	}

	for _, role := range roles {
		if val, ok := a.Roles[role]; ok && val {
			return true
		}
	}
	return false
}
