.PHONY: help run build clean test dev

# Variables
BINARY_NAME=htmx-live-search
PORT=8080

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	@echo "🚀 Starting server on http://localhost:$(PORT)"
	@go run main.go

build: ## Build the application
	@echo "🔨 Building binary..."
	@go build -o $(BINARY_NAME) main.go
	@echo "✅ Built: $(BINARY_NAME)"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f *.exe
	@echo "✅ Cleaned!"

test: ## Run tests (placeholder)
	@echo "🧪 Running tests..."
	@go test -v ./...

dev: ## Run with auto-reload (requires air)
	@echo "🔄 Running with auto-reload..."
	@air

init: ## Initialize go module
	@echo "📦 Initializing Go module..."
	@go mod init htmx-go-live-search
	@go mod tidy
	@echo "✅ Go module initialized!"

fmt: ## Format Go code
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted!"

lint: ## Run linter (requires golangci-lint)
	@echo "🔍 Running linter..."
	@golangci-lint run
	@echo "✅ Linting complete!"