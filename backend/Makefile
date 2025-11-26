.PHONY: help build run test clean migrate-up migrate-down docker-build docker-up docker-down

# Load .env file if it exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Variables
APP_NAME=ticketing-api
MAIN_PATH=cmd/api/main.go
BINARY_NAME=bin/$(APP_NAME)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BINARY_NAME)"

run: ## Run the application
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_PATH)

dev: ## Run with auto-reload (requires air)
	@echo "Running in development mode..."
	@air

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

mocks: ## Generate mocks using mockery
	@echo "Generating mocks..."
	@mockery --name=UserRepository --dir=internal/repositories --output=internal/repositories/mocks --outpkg=mocks
	@mockery --name=TicketRepository --dir=internal/repositories --output=internal/repositories/mocks --outpkg=mocks
	@mockery --name=EventRepository --dir=internal/repositories --output=internal/repositories/mocks --outpkg=mocks
	@echo "Mocks generated in internal/repositories/mocks/"

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

migrate-create: ## Create a new migration (usage: make migrate-create name=create_users_table)
	@echo "Creating migration: $(name)"
	@migrate create -ext sql -dir migrations -seq $(name)

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	@migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down: ## Rollback last migration
	@echo "Rolling back last migration..."
	@migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-drop: ## Drop all tables (WARNING: destructive)
	@echo "Dropping all tables..."
	@migrate -path migrations -database "$(DATABASE_URL)" drop -f

migrate-status: ## Show migration status
	@echo "Migration status:"
	@migrate -path migrations -database "$(DATABASE_URL)" version

migrate-force: ## Force set migration version (usage: make migrate-force version=1)
	@echo "Forcing migration version to: $(version)"
	@migrate -path migrations -database "$(DATABASE_URL)" force $(version)

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):latest .

docker-up: ## Start services with docker-compose
	@echo "Starting services..."
	@docker-compose up -d

docker-down: ## Stop services with docker-compose
	@echo "Stopping services..."
	@docker-compose down

docker-logs: ## View docker-compose logs
	@docker-compose logs -f

install-tools: ## Install development tools
	@echo "Installing tools..."
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/vektra/mockery/v2@latest
	@echo "Tools installed"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

.DEFAULT_GOAL := help