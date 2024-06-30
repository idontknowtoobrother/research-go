package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequirePermission(fctx *fiber.Ctx, permissionGroup map[string]bool) error {
	player := fctx.Locals("user").(*jwt.Token)
	claims := player.Claims.(jwt.MapClaims)

	playerGroup, ok := claims["group"].(string)
	if !ok {
		return fctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "permission group not found",
		})
	}

	if !permissionGroup[playerGroup] {
		return fctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "permission denied",
		})
	}

	return fctx.Next()
}
