# Makefile for a Go project

# Define the Go compiler
GO := go

# Define the project name
PROJECT_NAME := DiffieHellman

# Define the source code directory
SRC_DIR := ./

# Define the output directory
OUT_DIR := ./bin

# Define the main source file
MAIN := $(SRC_DIR)/main.go

# Define the test source files
TESTS := $(wildcard $(SRC_DIR)/*_test.go)

# Define the build target
build: $(MAIN)
	@echo "Building $(PROJECT_NAME)..."
	@$(GO) build -o $(OUT_DIR)/$(PROJECT_NAME).exe $(MAIN)

# Define the test target
test: $(TESTS)
	@echo "Running tests..."
	@$(GO) test -v $(SRC_DIR)test -run TestFindPrimitiveRoot

# Define the clean target
clean:
	@echo "Cleaning up..."
	@echo Remove-Item -Recurse $(OUT_DIR)
	@Remove-Item -Recurse $(OUT_DIR)

.PHONY: build test clean
