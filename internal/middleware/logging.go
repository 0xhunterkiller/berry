package middleware

import (
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func LogRequests(c *fiber.Ctx) error {
	logger.Logger.Info("got ", c.Method(), " request for: ", c.Path())
	return c.Next()
}
