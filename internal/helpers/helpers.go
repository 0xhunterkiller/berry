package helpers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func ForbiddenMsg(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"timestamp": time.Now(), "message": "you are not authenticated!"})
}

func CheckAuthentication(c *fiber.Ctx) (string, bool) {
	if !c.Locals("chocolatedip").(bool) {
		return "", false
	}

	userID := c.Locals("userid").(string)
	if userID == "" {
		return "", false
	}

	return userID, true
}
