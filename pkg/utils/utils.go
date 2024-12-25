package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironment loads environment variables from a .env file and verifies that required keys are set.
//
// This function uses the `godotenv` library to load variables from a .env file into the environment.
// If the .env file is missing, it logs a warning but continues execution.
//
// Parameters:
//   envkeys ...string: A variadic parameter of required environment variable keys.
//                      The function ensures that all provided keys are set in the environment.
//
// Behavior:
//   - If the .env file is not found, it prints a warning ("proper .env file not found").
//   - For each key in envkeys, it checks if the environment variable is set.
//   - If any key is missing, the program logs a fatal error and exits.
//
// Example:
//   LoadEnvironment("DATABASE_URL", "REDIS_URL")
//   // Exits the program if DATABASE_URL or REDIS_URL is not set.
//
// Dependencies:
//   - github.com/joho/godotenv for loading environment variables from a .env file.
//   - os.LookupEnv for checking environment variables.
//   - log.Fatalln for logging fatal errors and exiting.

func LoadEnvironment(envkeys ...string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("proper .env file not found")
	}

	for _, v := range envkeys {
		if _, ok := os.LookupEnv(v); !ok {
			log.Fatalln("Env is not set: ", v)
		}
	}
}
