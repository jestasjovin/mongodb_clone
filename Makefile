# Output Settings
BINARY_NAME=mongodb_clone
BUILD_DIR=dist

# Go-related settings
GO=go
GOTEST=$(GO) test
GOBUILD=$(GO) build

# Directories
SRC_DIR=internal/storage
TEST_DIR=tests

# Default target 
all: build

# Creating dist dir if  not exist
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

build: $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "Build complete."

# Run app
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up generated files
clean:
	rm -rf $(BUILD_DIR)
	@echo "Cleaned up."

# Testing the application
test:
	$(GOTEST) -v $(SRC_DIR)/*

# Run tests for  storage: pager/page functionality
test_storage:
	$(GOTEST) -v ./internal/storage/tests
	@echo "Pager and Page tests complete."

# Build the Go binary and run tests
build_and_test: build test

# Default target for quick run (build, test, and run app)
quick_run: build_and_test run

.PHONY: all build run clean test test_pager build_and_test quick_run
