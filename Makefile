# Project metadata
APP_NAME := berry
BUILD_DIR := bin
MAIN_PKG := ./cmd

# Default target
.PHONY: all
all: build

# Run the app
.PHONY: run
run:
	go run $(MAIN_PKG)

# Build the app
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PKG)

# Run tests
.PHONY: test
test:
	go test ./...

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Run linters (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Install/update dependencies
.PHONY: deps
deps:
	go mod tidy