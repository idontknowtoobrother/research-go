package main

import (
	"fmt"
	"os"
	"time"

	logCharmBracelet "github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/idontknowtoobrother/monolith-api-crud/models"
	"github.com/idontknowtoobrother/monolith-api-crud/services"
	"github.com/idontknowtoobrother/monolith-api-crud/stores"
)

var (
	log = logCharmBracelet.NewWithOptions(os.Stderr, logCharmBracelet.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "CRUD 🗞️",
	})
)

func main() {
	// Stores (Database)
	playerStore := stores.NewPlayerStore()

	// Services (Bussiness Logic)
	playerService := services.NewPlayerService(playerStore)

	// Fiber App
	app := fiber.New()

	// Routes
	app.Get("/connect", func(fctx *fiber.Ctx) error {
		return fctx.SendString("Joined, World! 🌍")
	})
	app.Get("/players", func(fctx *fiber.Ctx) error {
		return fctx.JSON(playerService.GetPlayers())
	})
	app.Get("/players/:uuid", func(fctx *fiber.Ctx) error {
		uuid := fctx.Params("uuid")
		player, err := playerService.GetPlayerByUuid(uuid)
		if err != nil {
			return fctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return fctx.JSON(player)
	})

	app.Post("/players", func(fctx *fiber.Ctx) error {
		var player models.Player
		if err := fctx.BodyParser(&player); err != nil {
			return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		createdPlayer, err := playerService.CreatePlayer(player)
		if err != nil {
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return fctx.JSON(createdPlayer)
	})

	app.Patch("/players", func(fctx *fiber.Ctx) error {
		var player models.Player
		if err := fctx.BodyParser(&player); err != nil {
			return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		updatedPlayer, err := playerService.UpdatePlayer(player)
		if err != nil {
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return fctx.JSON(updatedPlayer)
	})

	app.Delete("/players/:uuid", func(fctx *fiber.Ctx) error {
		uuid := fctx.Params("uuid")
		if err := playerService.DeletePlayerByUuid(uuid); err != nil {
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return fctx.JSON(fiber.Map{
			"message": fmt.Sprintf("Player uuid=%s deleted successfully", uuid),
		})
	})

	// Serve
	log.Info("Server is running on port 8000")
	app.Listen(":8000")
}
