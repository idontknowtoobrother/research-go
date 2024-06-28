package main

import (
	"net/http"
	"os"
	"time"

	"github.com/backend/middleware/core"
	"github.com/backend/middleware/handlers"
	"github.com/backend/middleware/middleware"
	logCharm "github.com/charmbracelet/log"
)

var (
	log = logCharm.NewWithOptions(os.Stderr, logCharm.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "Hex GATEWAY 👾 ",
	})
)

func main() {

	http.Handle("/connect", middleware.GetCredentials(middleware.Log(handlers.Connection{})))
	http.Handle("/credentials", middleware.ValidateCredentials(middleware.Log(handlers.Credentials{})))

	log.Info("server ready", "port", 8080)

	core.JobCredentials()

	http.ListenAndServe(":8080", nil)
}
