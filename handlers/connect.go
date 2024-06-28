package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/backend/middleware/middleware"
	logCharm "github.com/charmbracelet/log"
)

var (
	log = logCharm.NewWithOptions(os.Stderr, logCharm.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "Hex Connection 👾 ",
	})
)

type Connection struct{}

func (rcv Connection) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Info("connection established", "credential", r.Context().Value(middleware.CredentialKey("hex_credential")).(string))
	w.Write([]byte("connection established"))
}
