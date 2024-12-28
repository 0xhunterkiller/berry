package middleware

import (
	"github.com/0xhunterkiller/berry/pkg/jwtutil"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func unauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return unauthorizedResponse(c)
	}

	logger.Logger.Debug(tokenString)
	ok, token, err := jwtutil.CheckAndGetJWT(tokenString, []string{"HS256"}, "berry-authn", "userid")
	if !ok || err != nil {
		return unauthorizedResponse(c)
	}
	userID, ok := jwtutil.GetFromClaims(token, "userid")
	if !ok {
		return unauthorizedResponse(c)
	}
	c.Locals("userid", userID)
	return c.Next()
}
