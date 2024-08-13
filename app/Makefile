# Variables
BINARY_NAME := run-app
MAIN_PACKAGE := ./main.go

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Run the project
run: build
	@echo "Running the project..."
	./$(BINARY_NAME) "$(ARGS)"

# Test the project and generate coverage report
test:
	@echo "Running tests and generating coverage report..."
	mkdir -p test_coverage
	go test ./... -coverprofile=test_coverage/test.out
	go tool cover -html=./test_coverage/test.out -o test_coverage/coverage.html

# Clean up build files
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)
	rm -f test_coverage/test.out
	rm -rf test_coverage/coverage.html

# Phony targets to avoid conflicts with file names
.PHONY: all build run clean test
