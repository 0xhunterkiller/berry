package middleware

import (
	"strings"

	"github.com/0xhunterkiller/berry/pkg/jwtutil"
	"github.com/gofiber/fiber/v2"
)

func unauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return unauthorizedResponse(c)
	}

	authHeaderComps := strings.Split(authHeader, " ")
	if len(authHeaderComps) == 2 && authHeaderComps[0] == "Bearer" {
		tokenString := authHeaderComps[1]

		ok, token, err := jwtutil.CheckAndGetJWT(tokenString, []string{"HS256"}, "berry-authn", "userid")
		if !ok || err != nil {
			return unauthorizedResponse(c)
		}
		userID, ok := jwtutil.GetFromClaims(token, "userid")
		if !ok {
			return unauthorizedResponse(c)
		}
		c.Locals("userid", userID)
		c.Locals("chocolatedip", true)
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized, Auth Header Format is `Authorization: Bearer aabbxxyy1122`"})
}
