.PHONY: build run test clean install dev fmt lint vet deps

APP_NAME=bubble-skeleton
BUILD_DIR=build
CMD_DIR=cmd

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

run:
	@go run $(CMD_DIR)/main.go

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f debug.log
	@echo "Clean complete"

install:
	@echo "Installing $(APP_NAME)..."
	@go install $(CMD_DIR)/main.go
	@echo "Installation complete"

dev:
	@echo "Running in development mode with live reload..."
	@command -v air > /dev/null 2>&1 || (echo "Installing air..." && go install github.com/air-verse/air@latest)
	@air

fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

lint:
	@echo "Running linter..."
	@command -v golangci-lint > /dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@golangci-lint run ./...
	@echo "Lint complete"

vet:
	@echo "Running go vet..."
	@go vet ./...
	@echo "Vet complete"

deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies updated"

init: deps
	@echo "Initializing project..."
	@mkdir -p $(BUILD_DIR)
	@echo "Project initialized"

help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  test     - Run tests"
	@echo "  clean    - Remove build artifacts"
	@echo "  install  - Install the application"
	@echo "  dev      - Run with live reload (requires air)"
	@echo "  fmt      - Format code"
	@echo "  lint     - Run linter (requires golangci-lint)"
	@echo "  vet      - Run go vet"
	@echo "  deps     - Download and tidy dependencies"
	@echo "  init     - Initialize project"
	@echo "  help     - Show this help message"