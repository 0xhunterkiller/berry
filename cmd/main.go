package main

import (
	"log"
	"time"

	"github.com/0xhunterkiller/berry/internal/store"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/0xhunterkiller/berry/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// load env vars
	utils.LoadEnvironment("LOG_LEVEL", "JWT_KEY", "PSQL_HOST", "PSQL_PORT", "PSQL_USER", "PSQL_PASSWORD", "PSQL_DB", "PSQL_SSLMODE", "MIG_DIR")

	// initialize logger
	logger.InitLogger()

	// initialize store
	store.InitStore()

	// start fiber app
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"timestamp": time.Now()})
	})

	log.Fatalln(app.Listen(":3000"))
}
