package middleware

import (
	"errors"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"

	"github.com/LazyCodeTeam/just-code-backend/internal/api/util"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/model"
	"github.com/LazyCodeTeam/just-code-backend/internal/core/usecase"
	coreUtil "github.com/LazyCodeTeam/just-code-backend/internal/core/util"
)

type AuthTokenValidator struct {
	client *auth.Client
}

func NewAuthTokenValidator(client *auth.Client) *AuthTokenValidator {
	return &AuthTokenValidator{
		client: client,
	}
}

func (m *AuthTokenValidator) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)

		result, err := m.client.VerifyIDToken(r.Context(), token)
		if err != nil {
			util.WriteError(w, model.NewError(usecase.ErrorTypeUnauthorized))
			return
		}
		authData, err := getAuthDataFromToken(result)
		if err != nil {
			util.WriteError(w, model.NewError(usecase.ErrorTypeUnauthorized))
			return
		}
		ctx := coreUtil.ContextWithAuthData(r.Context(), authData)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
	return nil, errors.New("unknown auth type")
}

func getEmailAuthDataFromToken(token *auth.Token) (*model.AuthData, error) {
	email, ok := token.Claims["email"].(string)
	if !ok {
		return nil, errors.New("email not found")
	}
	verified, ok := token.Claims["email_verified"].(bool)
	if !ok {
		return nil, errors.New("email verified not found")
	}

	return &model.AuthData{
		Type:     model.AuthTypeEmail,
		Email:    &email,
		Verified: verified,
		Id:       token.UID,
	}, nil
}
