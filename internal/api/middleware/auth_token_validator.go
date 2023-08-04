package middleware

import (
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
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

		_, err := m.client.VerifyIDToken(r.Context(), token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
