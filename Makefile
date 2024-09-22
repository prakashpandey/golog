# Variables
MODULE_NAME = golog
PKG = ./...
OUTPUT_DIR = ./dist
VERSION = $(shell git describe --tags --always)
GO_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Targets
.PHONY: all clean test release

all: test

test:
	@echo "Running tests with coverage..."
	go test $(PKG) -cover -coverprofile=coverage.out
	go tool cover -func=coverage.out

build:
	@echo "Building the module (no binary)..."
	go build $(PKG)

release: clean
	@echo "Creating release artifact..."
	mkdir -p $(OUTPUT_DIR)
	tar -czvf $(OUTPUT_DIR)/$(MODULE_NAME)-$(VERSION).tar.gz $(GO_FILES) README.md LICENSE go.mod go.sum

clean:
	@echo "Cleaning up..."
	rm -rf $(OUTPUT_DIR)
	rm -f coverage.out
