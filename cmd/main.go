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
	"github.com/jmoiron/sqlx"
)

func gracefulShutdown(app *fiber.App, db *sqlx.DB) {
	if app != nil {
		err := app.Shutdown()
		if err != nil {
			logger.Logger.Error("Encountered an error, while shutting down ", err)
		}
	}
	if db != nil {
		dbpsql.CloseDBConn(db)
	}
	os.Exit(1)
}

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
		"MIG_DIR",
		"ADMIN_PASSWORD"}
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

	var app *fiber.App
	var db *sqlx.DB

	db, err := dbpsql.ConnectDB(psqlInfo, 10, 5, 30)
	if err != nil {
		logger.Logger.Errorf("Failed to connect to the database: %v", err)
		gracefulShutdown(app, db)
	}

	if err := db.Ping(); err != nil {
		logger.Logger.Errorf("failed to connect to the database: %v", err)
		gracefulShutdown(app, db)
	}

	defer dbpsql.CloseDBConn(db)

	// Migration Up
	dbpsql.MigUp(db)

	adminUserID, err := appinit.CreateAdminUser(db, "admin@example.com", os.Getenv("ADMIN_PASSWORD"))
	if err == nil {
		if adminUserID != "" {
			logger.Logger.Info("berryroot was successfully created")
		}
	} else {
		logger.Logger.Errorf("berryroot was not created: %v", err.Error())
		gracefulShutdown(app, db)
	}

	adminRoleID, err := appinit.CreateRootRole(db)
	if err == nil {
		if adminRoleID != "" {
			logger.Logger.Info("root role was successfully created")
		}
	} else {
		logger.Logger.Errorf("root role was not created: %v", err.Error())
		gracefulShutdown(app, db)
	}

	err = appinit.MakeRoot(db, adminUserID, adminRoleID)
	if err != nil {
		logger.Logger.Errorf("berryroot is not root, role couldn't be assigned: %v", err.Error())
	}
	logger.Logger.Info("berryroot is now root")

	// Prepare an injection
	inj := &models.Deps{DB: db}

	// Initialize Application Layers (Handlers -> Services -> Store)
	handlers := appinit.AppInit(inj)

	// Start Fiber app
	app = fiber.New()
	app.Use(middleware.LogRequests)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"timestamp": time.Now()})
	})

	// Register User Handler
	uh := handlers.UserHandler
	uh.RegisterRoutes(app)

	rh := handlers.RoleHandler
	rh.RegisterRoutes(app)

	ah := handlers.ActionHandler
	ah.RegisterRoutes(app)

	ph := handlers.PermissionHandler
	ph.RegisterRoutes(app)

	reh := handlers.ResourceHandler
	reh.RegisterRoutes(app)

	man := handlers.ManagerHandler
	man.RegisterRoutes(app)

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		logger.Logger.Info("Gracefully shutting down...")
		gracefulShutdown(app, db)
	}()

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}
