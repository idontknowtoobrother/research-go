package handlers

import (
	"net/http"

	"github.com/backend/middleware/middleware"
)

type Credentials struct{}

func (rcv Credentials) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("credentials validated", "credential", r.Context().Value(middleware.CredentialKey("hex_credential")))
	w.Write([]byte("credentials validated"))
}
