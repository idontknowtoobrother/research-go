package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/idontknowtoobrother/monolith-api-crud-common/utils"
	"github.com/idontknowtoobrother/monolith-api-crud/middlewares"
	"github.com/idontknowtoobrother/monolith-api-crud/services"
	"github.com/idontknowtoobrother/monolith-api-crud/stores"

	logCharmBracelet "github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
)

var (
	log = logCharmBracelet.NewWithOptions(os.Stderr, logCharmBracelet.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "CRUD 🗞️",
	})
)

var (
	serverPort = utils.GetEnv(".game-server.env", "SERVER_PORT", ":3000")
)

func main() {
	// Stores (Database)
	playerStore := stores.NewPlayerStore()

	// Services (Bussiness Logic)
	playerService := services.NewPlayerService(playerStore)

	// Fiber App
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOriginsFunc: func(origin string) bool {
			log.Info("Origin: ", origin)
			return true
		},
	}))

	app.Use(middlewares.Log)

	// Routes
	InitRoutes(app, playerStore, playerService)

	// Serve
	log.Info("Server is running on port ", "port", serverPort)
	app.Listen(serverPort)
}
