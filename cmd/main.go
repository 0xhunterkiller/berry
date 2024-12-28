package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xhunterkiller/berry/internal/appinit"
	"github.com/0xhunterkiller/berry/internal/middleware"
	"github.com/0xhunterkiller/berry/internal/models"
	"github.com/0xhunterkiller/berry/pkg/dbpsql"
	"github.com/0xhunterkiller/berry/pkg/logger"
	"github.com/0xhunterkiller/berry/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// load env vars
	envKeys := []string{
		"LOG_LEVEL",
		"APP_PORT",
		"JWT_KEY",
		"PSQL_HOST",
		"PSQL_PORT",
		"PSQL_USER",
		"PSQL_PASSWORD",
		"PSQL_DB",
		"PSQL_SSLMODE",
		"MIG_DIR"}
	utils.LoadEnvironment(envKeys...)

	// initialize logger
	logger.InitLogger()

	// Open DB Connection
	psqlInfo := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DB"),
		os.Getenv("PSQL_SSLMODE"))

	db, err := dbpsql.ConnectDB(psqlInfo, 10, 5, 30)
	if err != nil {
		logger.Logger.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		logger.Logger.Fatalf("failed to connect to the database: %v", err)
	}

	defer dbpsql.CloseDBConn(db)

	// Migration Up
	dbpsql.MigUp(db)

	// Prepare an injection
	inj := &models.Deps{DB: db}

	// Initialize Application Layers (Handlers -> Services -> Store)
	handlers := appinit.AppInit(inj)

	// Start Fiber app
	app := fiber.New()
	app.Use(middleware.LogRequests)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"timestamp": time.Now()})
	})

	// Register User Handler
	uh := handlers.UserHandler
	uh.RegisterRoutes(app)

	// Graceful Shutdown
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
