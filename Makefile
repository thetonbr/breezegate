# Variables
APP_NAME := breezegate
BUILD_DIR := ./build
CMD_DIR := ./cmd/app
BINARY_PATH := $(BUILD_DIR)/$(APP_NAME)
GO_FILES := $(shell find . -name '*.go')

# Go commands
GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_TEST := $(GO) test
GO_FMT := $(GO) fmt
GO_LINT := golangci-lint run
GO_DOC := godoc
GO_MOD := $(GO) mod
GO_CLEAN := $(GO) clean
GO_COVERAGE := $(GO) tool cover

# Test directory
TEST_DIR := ./test

# Directories for the source code and test files
SRC_DIRS := ./internal/domain ./internal/services ./cmd/app

# Default target: build
all: build

# Build the application
build:
	@echo "Building the application..."
	$(GO_BUILD) -o $(BINARY_PATH) $(CMD_DIR)

# Run the application
run: build
	@echo "Running the application..."
	$(BINARY_PATH)

# Run tests with coverage
test:
	@echo "Running tests with coverage..."
	$(GO_TEST) -cover -coverprofile=coverage.out $(SRC_DIRS) ./test/

# Display coverage report in the terminal
coverage:
	@echo "Showing test coverage..."
	$(GO_COVERAGE) -func=coverage.out

# Generate HTML coverage report
coverage-html:
	@echo "Generating HTML coverage report..."
	$(GO_COVERAGE) -html=coverage.out -o coverage.html
	@echo "HTML coverage report generated as coverage.html"

# Format code
fmt:
	@echo "Formatting Go code..."
	$(GO_FMT) ./...

# Lint the code
lint:
	@echo "Linting Go code..."
	$(GO_LINT)

# Clean build artifacts
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
	$(GO_CLEAN)


# Clean test artifacts
clean-tests:
	@echo "Cleaning test artifacts..."
	rm -f coverage.out coverage.html

# Generate Go documentation in HTML
doc:
	@echo "Generating Go documentation..."
	$(GO_DOC) -http=:6060 &
	@echo "Go documentation is being served at http://localhost:6060"

# Generate Go documentation as Markdown
gendoc:
	@echo "Generating Go documentation as Markdown..."
	$(GO) install github.com/robertkrimen/godocdown/godocdown@latest
	@echo "Generating Markdown documentation for the project..."
	godocdown -o DOC.md
	@echo "Documentation has been generated and written to DOC.md"


# Generate Go documentation for the package
docpkg:
	@echo "Generating Go package documentation..."
	$(GO_DOC) -all > package_doc.txt
	@echo "Go documentation has been written to package_doc.txt"

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .

docker-run:
	@echo "Running Docker container..."
	docker run -p 80:80 -p 443:443 --name $(APP_NAME) $(APP_NAME):latest

# clean:
# 	@echo "Cleaning test artifacts..."
# 	rm -f coverage.out

# Default clean task
.PHONY: clean
