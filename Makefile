# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=echotonic
MAIN_PATH=.
GO_PATH=$(shell go env GOPATH)

# Show help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  run           - Run the application"
	@echo "  build         - Build the application"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  clean         - Clean build files"
	@echo "  deps          - Download dependencies"
	@echo "  tidy          - Tidy up dependencies"
	@echo "  update        - Update dependencies"
	@echo "  dev           - Run with hot reload (requires air)"
	@echo "  install-tools - Install development tools"
	@echo "  help          - Show this help"

# Run the application
.PHONY: run
run:
	$(GOCMD) run $(MAIN_PATH)

# Build the application
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

# Test the application
.PHONY: test
test:
	$(GOTEST) -v ./...

# Test with coverage
.PHONY: test-coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Clean build files
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

# Download dependencies
.PHONY: deps
deps:
	$(GOMOD) download
	$(GOMOD) verify

# Tidy up dependencies
.PHONY: tidy
tidy:
	$(GOMOD) tidy

# Update dependencies
.PHONY: update
update:
	$(GOGET) -u ./...
	$(GOMOD) tidy

# Run development server with hot reload (requires air to be installed)
.PHONY: dev
dev:
	${GO_PATH}/bin/air -c .air.toml

.PHONY: install-tools
install-tools:
	$(GOGET) github.com/air-verse/air@latest

.PHONY: version
version:
	sed -i.bck "s|SERVICE_VERSION=\"[0-9]*.[0-9]*.[0-9]*.*\"|SERVICE_VERSION=\"${VERSION}\"|" "Dockerfile"
	rm -fr "Dockerfile.bck"
	git add "Dockerfile"
	git commit -m "bump v${VERSION}"
	git tag v${VERSION}