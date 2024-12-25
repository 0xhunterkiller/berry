# Variables
APP_NAME := berry
BUILD_DIR := bin

# Default target
all: build

# Build the Go binary
build:
	@echo "Building the application..."
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/main.go

# Run lint
lint:
	@echo "Running linter..."
	golangci-lint run

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Run the application
run: build
	@echo "Running the application..."
	./$(BUILD_DIR)/$(APP_NAME)

.PHONY: all build lint clean run
