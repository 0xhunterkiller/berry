package main

import (
	"log"
	"time"

	"github.com/0xhunterkiller/berry/internal/appinit"
	"github.com/0xhunterkiller/berry/pkg/dbpsql"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/0xhunterkiller/berry/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// load env vars
	utils.LoadEnvironment("LOG_LEVEL", "JWT_KEY", "PSQL_HOST", "PSQL_PORT", "PSQL_USER", "PSQL_PASSWORD", "PSQL_DB", "PSQL_SSLMODE", "MIG_DIR")

	// initialize logger
	logger.InitLogger()

	// Migration Up
	db, err := dbpsql.ConnectDB(10, 5, 30)
	if err != nil {
		logger.Logger.Errorf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		logger.Logger.Errorf("failed to connect to the database: %v", err)
	}

	defer dbpsql.CloseDBConn(db)
	dbpsql.MigUp(db)

	// Initialize Application
	handlers := appinit.AppInit(db)

	// start fiber app
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"timestamp": time.Now()})
	})

	// Register Auth Handler
	ah := handlers.AuthHandler
	ah.RegisterRoutes(app)

	log.Fatalln(app.Listen(":3000"))
}
