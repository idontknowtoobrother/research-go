package middleware

import (
	"context"
	"net/http"

	"github.com/backend/middleware/core"
)

type CredentialKey string

func GetCredentials(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credential := core.NewCredentials()
		ctx := context.WithValue(r.Context(), CredentialKey("hex_credential"), credential)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateCredentials(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credential := r.Context().Value(CredentialKey("hex_credential"))
		if credential == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		credentialString := credential.(string)
		if credentialString == "" {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		if !core.ValidateCredentials(credentialString) {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
