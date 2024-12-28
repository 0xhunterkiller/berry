package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// InitLogger initializes a global logger instance using the logrus library.
//
// This function creates a new logrus.Logger instance and configures it with a JSON formatter.
// It sets the log level based on the LOG_LEVEL environment variable. If LOG_LEVEL is not set or invalid,
// it defaults to the "info" log level.
//
// Behavior:
//   - The logger is assigned to the global `Logger` variable.
//   - The log format is set to JSON using `logrus.JSONFormatter`.
//   - The log level is determined by the LOG_LEVEL environment variable (case-insensitive).
//   - Supported levels: "panic", "fatal", "error", "warn", "info", "debug", "trace".
//   - If LOG_LEVEL is not set or an invalid level is provided, the log level defaults to "info".
//   - Prints a message to the console if an invalid log level is specified.
//
// Example:
//
//	InitLogger()
//	Logger.Info("Logger initialized successfully")
//
// Dependencies:
//   - github.com/sirupsen/logrus for logging.
//   - os.Getenv for retrieving the LOG_LEVEL environment variable.
//   - strings.ToLower for case-insensitive log level parsing.
//
// Global Variables:
//   - Logger: A pointer to the logrus.Logger instance, accessible globally.
//
// Notes:
//   - Ensure the LOG_LEVEL environment variable is set correctly in your environment.
//   - Typical log levels include "info", "warn", and "error".
var Logger *logrus.Logger

func InitLogger() *logrus.Logger {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info" // Default to "info" if LOG_LEVEL is not set
	}

	// Set the log level
	level, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		fmt.Printf("Invalid log level, defaulting to INFO: %v\n", err)
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)
	return Logger
}
