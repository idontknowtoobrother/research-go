package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/idontknowtoobrother/monolith-api-crud-common/utils"
	"github.com/idontknowtoobrother/monolith-api-crud/middlewares"
	"github.com/idontknowtoobrother/monolith-api-crud/models"
	"github.com/idontknowtoobrother/monolith-api-crud/services"
	"github.com/idontknowtoobrother/monolith-api-crud/stores"

	jwtware "github.com/gofiber/contrib/jwt"
)

type Credential struct {
	Key string `json:"key"`
}

func connectWorld(fctx *fiber.Ctx) error {
	var credential Credential
	if err := fctx.BodyParser(&credential); err != nil {
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if credential.Key != "SKHJIDCU@#$*&.TASUDJUWWIDDK" {
		return fiber.ErrUnauthorized
	}

	claims := jwt.MapClaims{
		"iss":   "game-server",
		"group": "user",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	secret := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	responseSecret, err := secret.SignedString([]byte(utils.GetEnv(".game-server", "JWT_SECRET_KEY", "default_secret")))
	if err != nil {
		return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return fctx.JSON(fiber.Map{
		"secret":  responseSecret,
		"message": "Joined, World! 🌍",
	})
}

func InitRoutes(fiberApp *fiber.App, playerStore *stores.PlayerStore, playerService *services.PlayerService) {

	// Routes (Public)
	fiberApp.Get("/connect", connectWorld)

	jwtConfig := jwtware.Config{
		AuthScheme: "Bearer",
		SigningKey: jwtware.SigningKey{Key: []byte(utils.GetEnv(".game-server", "JWT_SECRET_KEY", "default_secret"))},
		ErrorHandler: func(fctx *fiber.Ctx, err error) error {
			return fctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}

	playerGroup := fiberApp.Group("/players")
	playerGroup.Use(middlewares.Log)
	playerGroup.Use(jwtware.New(jwtConfig))

	// Routes (Protected)
	playerGroup.Get("/", func(fctx *fiber.Ctx) error {
		return fctx.JSON(playerService.GetPlayers())
	})

	playerGroup.Get("/:uuid", func(fctx *fiber.Ctx) error {
		uuid := fctx.Params("uuid")
		player, err := playerService.GetPlayerByUuid(uuid)
		if err != nil {
			return fctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return fctx.JSON(player)
	})

	playerGroup.Post("/", func(fctx *fiber.Ctx) error {
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

	playerGroup.Patch("/", func(fctx *fiber.Ctx) error {
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

	// Middleware (Protected) Admin Only
	playerGroup.Use(func(fctx *fiber.Ctx) error {
		requirePermissionGroup := map[string]bool{
			"admin": true,
		}
		return middlewares.RequirePermission(fctx, requirePermissionGroup)
	})

	playerGroup.Delete("/:uuid", func(fctx *fiber.Ctx) error {
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
}
