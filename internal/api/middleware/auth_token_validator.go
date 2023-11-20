package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/failure"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	coreUtil "github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type AuthTokenValidatorFactory struct {
	client *auth.Client
}

func NewAuthTokenValidator(client *auth.Client) *AuthTokenValidatorFactory {
	return &AuthTokenValidatorFactory{
		client: client,
	}
}

func (m *AuthTokenValidatorFactory) Get(
	allowedRoles ...string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			token = strings.Replace(token, "Bearer ", "", 1)

			result, err := m.client.VerifyIDToken(r.Context(), token)
			if err != nil {
				e := failure.NewAuthFailure(failure.FailureTypeUnauthorized, err)
				util.WriteError(w, e)
				return
			}
			authData, err := getAuthDataFromToken(result)
			if err != nil {
				slog.WarnContext(r.Context(), "Error getting auth data from token", "err", err)
				e := failure.NewAuthFailure(failure.FailureTypeUnauthorized, err)
				util.WriteError(w, e)
				return
			}
			ctx := coreUtil.ContextWithAuthData(r.Context(), authData)

			if !authData.CanAccess(allowedRoles...) {
				slog.WarnContext(
					ctx,
					"User is not allowed to access this resource",
					"authData",
					authData,
					"allowedRoles",
					allowedRoles,
				)
				util.WriteError(w, failure.NewNotFoundFailure(failure.FailureTypeNotFound))
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getAuthDataFromToken(token *auth.Token) (*model.AuthData, error) {
	switch token.Firebase.SignInProvider {
	case "anonymous":
		return &model.AuthData{
			Type: model.AuthTypeAnonymous,
			Id:   token.UID,
		}, nil
	case "password":
		return getEmailAuthDataFromToken(token)
	}
	return nil, fmt.Errorf("Unknown sign in provider: %s", token.Firebase.SignInProvider)
}

func getEmailAuthDataFromToken(token *auth.Token) (*model.AuthData, error) {
	email, ok := token.Claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("Email not found in token: %v", token.Claims)
	}
	verified, ok := token.Claims["email_verified"].(bool)
	if !ok {
		return nil, fmt.Errorf("Email verified not found in token: %v", token.Claims)
	}

	return &model.AuthData{
		Type:     model.AuthTypeEmail,
		Email:    &email,
		Verified: verified,
		Id:       token.UID,
		Roles:    getRolesFromToken(token),
	}, nil
}

func getRolesFromToken(token *auth.Token) map[string]bool {
	roles := map[string]bool{}
	if isAdmin, ok := token.Claims["admin"].(bool); ok && isAdmin {
		roles[model.AuthRoleAdmin] = true
	}

	return roles
}
