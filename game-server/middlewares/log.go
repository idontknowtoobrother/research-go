package middlewares

import (
	"os"

	logCharmBracelet "github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

var (
	log = logCharmBracelet.NewWithOptions(os.Stderr, logCharmBracelet.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      "15:04:05",
		Prefix:          "Log 🕵🏻",
	})
)

func Log(fctx *fiber.Ctx) error {
	log.Info("", "method", fctx.Method(), "path", fctx.Path(), "ip", fctx.IP())
	return fctx.Next()
}
