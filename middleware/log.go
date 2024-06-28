package middleware

import (
	"net/http"
	"os"
	"time"

	charmLog "github.com/charmbracelet/log"
)

var (
	log = charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "Hex Logs 👾 ",
	})
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info("request received", "credential", r.Context().Value(CredentialKey("hex_credential")), "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
