package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xhunterkiller/berry/internal/appinit"
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/pkg/dbpsql"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/0xhunterkiller/berry/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// load env vars
	utils.LoadEnvironment("LOG_LEVEL", "APP_PORT", "JWT_KEY", "PSQL_HOST", "PSQL_PORT", "PSQL_USER", "PSQL_PASSWORD", "PSQL_DB", "PSQL_SSLMODE", "MIG_DIR")

	// initialize logger
	l := logger.InitLogger()

	// Migration Up
	db, err := dbpsql.ConnectDB(10, 5, 30)
	if err != nil {
		l.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		l.Fatalf("failed to connect to the database: %v", err)
	}

	defer dbpsql.CloseDBConn(db)
	dbpsql.MigUp(db)

	ston := &models.Deps{DB: db, Logger: l}

	// Initialize Application
	handlers := appinit.AppInit(ston)

	// start fiber app
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"timestamp": time.Now()})
	})

	// Register Auth Handler
	ah := handlers.AuthHandler
	ah.RegisterRoutes(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		err = app.Shutdown()
		if err != nil {
			log.Println("Encountered an error, while shutting down ", err)
		}
		dbpsql.CloseDBConn(db)
	}()

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}
