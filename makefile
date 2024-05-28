# Makefile for student_scheduler project

# Variables
APP_NAME := student_scheduler
PORT := 3000

# Go commands
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOCLEAN := $(GOCMD) clean

# Default target
all: build

# Build the binary
build:
	@echo "Building the application..."
	$(GOBUILD) -o $(APP_NAME) ./main.go

# Run the application
run:
	@echo "Running the application..."
	$(GORUN) ./main.go

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOGET) -v -t -d ./...

# Clean the build
clean:
	@echo "Cleaning the build..."
	$(GOCLEAN)
	@rm -f $(APP_NAME)

# Start the server
start:
	@echo "Starting the server..."
	$(GORUN) main.go

# Stop the server (placeholder for actual stop command)
stop:
	@echo "Stopping the server..."

# Setup environment variables (example usage)
env:
	@echo "Setting up environment variables..."
	export PORT=$(PORT)

# PHONY targets to avoid filename conflicts
.PHONY: all build run test deps clean start stop env
