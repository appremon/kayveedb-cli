# Variables
APP_NAME := kayveedb-cli
TOOLS_DIR := tools
BIN_DIR := bin
VERSION_FILE := VERSION
VERSION_GO := cmd/version.go
VERSION_TEST_GO := cmd/version_test.go
README_FILE := README.md

# Read the current version from the VERSION file
CURRENT_VERSION := $(shell cat $(VERSION_FILE))
VERSION_MAJOR := $(word 1, $(subst ., ,$(CURRENT_VERSION)))
VERSION_MINOR := $(word 2, $(subst ., ,$(CURRENT_VERSION)))
VERSION_PATCH := $(word 3, $(subst ., ,$(CURRENT_VERSION)))

# Default target
all: build build-tools

# Build the main application
build: 
	@echo "Building $(APP_NAME)..."
	@go build -o $(BIN_DIR)/$(APP_NAME) main.go
	@echo "$(APP_NAME) built successfully."

# Traverse the tools directory and build each tool
build-tools:
	@echo "Building tools..."
	@mkdir -p $(BIN_DIR)
	@find $(TOOLS_DIR) -type f -name "main.go" | while read -r tool; do \
		TOOL_NAME=$$(dirname $$tool | xargs basename); \
		echo "Building $$TOOL_NAME..."; \
		go build -o $(BIN_DIR)/$$TOOL_NAME $$tool; \
		echo "$$TOOL_NAME built successfully."; \
	done

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean up binaries
clean:
	@echo "Cleaning up binaries..."
	@rm -rf $(BIN_DIR)
	@echo "Cleaned up successfully."

# Lint the code (requires golangci-lint to be installed)
lint:
	@echo "Linting code..."
	@~/go/bin/golangci-lint run

# Push version to GitHub (for use with your pushversion.sh script)
push:
	@echo "Pushing new version to GitHub..."
	@./githubBuild/pushversion.sh

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download

# Help message
help:
	@echo "Makefile commands:"
	@echo "  make build            - Build the main application"
	@echo "  make test             - Run tests"
	@echo "  make clean            - Clean up binaries"
	@echo "  make lint             - Lint the code"
	@echo "  make release          - Push new version to GitHub"
	@echo "  make deps             - Install dependencies"
	@echo "  make increment-minor  - Increment the minor version number"
	@echo "  make increment-major  - Increment the major version number"
	@echo "  make increment-patch  - Increment the patch version number"

check-versions:
	@echo "Checking version consistency..."
	@VERSION=$(shell cat $(VERSION_FILE) | tr -d '[:space:]') && \
	echo "VERSION file: $$VERSION" && \
	VERSION_GO_FILE=$$(grep 'const Version string' $(VERSION_GO) | grep -o 'v[0-9]\+\.[0-9]\+\.[0-9]\+\b' | tr -d '[:space:]') && \
	echo "cmd/version.go: $$VERSION_GO_FILE" && \
	VERSION_TEST_FILE=$$(grep 'expected := "kvdbcli version:' $(VERSION_TEST_GO) | grep -o 'v[0-9]\+\.[0-9]\+\.[0-9]\+\b' | tr -d '[:space:]') && \
	echo "cmd/version_test.go: $$VERSION_TEST_FILE" && \
	README_VERSION=$$(grep -o 'Current version: \*\*v[0-9]\+\.[0-9]\+\.[0-9]\+\*\*' $(README_FILE) | grep -o 'v[0-9]\+\.[0-9]\+\.[0-9]\+' | tr -d '[:space:]') && \
	echo "README.md: $$README_VERSION"

# Update the version in kayveedb.go and version_test.go
update-version-go:
	@echo "Updating version in $(VERSION_GO)"
	@sed -i 's/const Version string = "v[0-9]\+\.[0-9]\+\.[0-9]\+"/const Version string = "v$(NEW_VERSION)"/' $(VERSION_GO)
	@echo "Updated version in $(VERSION_GO)"
	@echo "Updating version in $(VERSION_TEST_GO)"
	@sed -i 's/expected := "kvdbcli version: v[0-9]\+\.[0-9]\+\.[0-9]\+\\n"/expected := "kvdbcli version: v$(NEW_VERSION)\\n"/' $(VERSION_TEST_GO)
	@echo "Updated version in $(VERSION_TEST_GO)"

# Update the version in README.md
update-version-readme:
	@echo "Updating version in $(README_FILE)"
	@sed -i 's/Current version: \*\*v[0-9]\+\.[0-9]\+\.[0-9]\+\*\*/Current version: \*\*v$(NEW_VERSION)\*\*/' $(README_FILE)
	@echo "Updated version in $(README_FILE)"

# Increment version numbers
increment-patch:
	@NEW_VERSION=$(VERSION_MAJOR).$(VERSION_MINOR).$$(( $(VERSION_PATCH) + 1 )) && \
	echo "Updating VERSION file from $(CURRENT_VERSION) to $$NEW_VERSION" && \
	echo $$NEW_VERSION > $(VERSION_FILE) && \
	$(MAKE) update-version-go NEW_VERSION=$$NEW_VERSION && \
	$(MAKE) update-version-readme NEW_VERSION=$$NEW_VERSION && \
	echo "Version updated to $$NEW_VERSION."

increment-minor:
	@NEW_VERSION=$(VERSION_MAJOR).$$(( $(VERSION_MINOR) + 1 )).0 && \
	echo "Updating VERSION file from $(CURRENT_VERSION) to $$NEW_VERSION" && \
	echo $$NEW_VERSION > $(VERSION_FILE) && \
	$(MAKE) update-version-go NEW_VERSION=$$NEW_VERSION && \
	$(MAKE) update-version-readme NEW_VERSION=$$NEW_VERSION && \
	echo "Version updated to $$NEW_VERSION."

increment-major:
	@NEW_VERSION=$$(( $(VERSION_MAJOR) + 1 )).0.0 && \
	echo "Updating VERSION file from $(CURRENT_VERSION) to $$NEW_VERSION" && \
	echo $$NEW_VERSION > $(VERSION_FILE) && \
	$(MAKE) update-version-go NEW_VERSION=$$NEW_VERSION && \
	$(MAKE) update-version-readme NEW_VERSION=$$NEW_VERSION && \
	echo "Version updated to $$NEW_VERSION."

# Push version to GitHub (for use with your pushversion.sh script)
release: 
	@echo "Pushing new version to GitHub..."
	@./githubBuild/pushversion.sh

.PHONY: all build test clean lint release deps help increment-patch increment-minor increment-major
