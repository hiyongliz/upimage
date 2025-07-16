# Project variables
BINARY_NAME=upimage
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date +%Y-%m-%dT%H:%M:%S%z)
COMMIT_SHA=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Go variables
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
CGO_ENABLED=0
GO111MODULE=on

# Build flags
LDFLAGS=-ldflags "-s -w -X upimage/cmd.Version=$(VERSION) -X upimage/cmd.BuildTime=$(BUILD_TIME) -X upimage/cmd.CommitSHA=$(COMMIT_SHA)"

# Default target
.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build the binary for current platform
	CGO_ENABLED=$(CGO_ENABLED) GO111MODULE=$(GO111MODULE) go build $(LDFLAGS) -o $(BINARY_NAME) main.go

.PHONY: bin
bin: build ## Alias for build (backward compatibility)

.PHONY: install
install: build ## Install the binary to GOPATH/bin
	go install $(LDFLAGS) .

.PHONY: clean
clean: ## Remove built binaries and cache
	rm -f $(BINARY_NAME)
	rm -rf dist/
	go clean -cache

.PHONY: test
test: ## Run all tests
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: lint
lint: ## Run linter (requires golangci-lint)
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run

.PHONY: fmt
fmt: ## Format Go code
	go fmt ./...
	go mod tidy

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: check
check: fmt vet lint test ## Run all checks (format, vet, lint, test)

# Cross-compilation targets
.PHONY: build-linux
build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 main.go

.PHONY: build-windows
build-windows: ## Build for Windows
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe main.go

.PHONY: build-darwin
build-darwin: ## Build for macOS
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 main.go

.PHONY: build-darwin-arm64
build-darwin-arm64: ## Build for macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 main.go

.PHONY: build-all
build-all: build-linux build-windows build-darwin build-darwin-arm64 ## Build for all platforms

.PHONY: release
release: clean check build-all ## Prepare a release (clean, check, and build for all platforms)
	@echo "Release built successfully!"
	@echo "Binaries available in dist/ directory:"
	@ls -la dist/

.PHONY: deps
deps: ## Download and verify dependencies
	go mod download
	go mod verify

.PHONY: deps-update
deps-update: ## Update dependencies
	go get -u ./...
	go mod tidy

.PHONY: run
run: ## Run the application with example arguments
	go run main.go --help

.PHONY: docker-build
docker-build: ## Build Docker image (if Dockerfile exists)
	@if [ -f Dockerfile ]; then \
		docker build -t $(BINARY_NAME):$(VERSION) .; \
	else \
		echo "Dockerfile not found"; \
	fi

.PHONY: version
version: ## Show version information
	@echo "Binary: $(BINARY_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit: $(COMMIT_SHA)"
	@echo "Go Version: $(shell go version)"
