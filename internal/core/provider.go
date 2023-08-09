package core

import "github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"

func Providers() []interface{} {
	return []interface{}{
		usecase.NewGetCurrentUser,
	}
}