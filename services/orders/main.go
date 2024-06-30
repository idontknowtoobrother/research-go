package main

import (
	"os"
	"time"

	logCharmBracelet "github.com/charmbracelet/log"
)

var (
	log = logCharmBracelet.NewWithOptions(os.Stderr, logCharmBracelet.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "Orders 🍣",
	})
)

func main() {
	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()
}
