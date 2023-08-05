package util

import (
	"context"

	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
)

const authDataKey = "authData"

func ContextWithAuthData(ctx context.Context, authData *model.AuthData) context.Context {
	return context.WithValue(ctx, authDataKey, authData)
}

func ExtractAuthData(ctx context.Context) *model.AuthData {
	authData, ok := ctx.Value(authDataKey).(*model.AuthData)
	if !ok {
		return nil
	}
	return authData
}
