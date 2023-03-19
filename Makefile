# Set the project name
PROJECT_NAME=shorty

APP_PORT=8080
REDIS_PORT=6379

.PHONY: build
## Build the Go application
build:
	@echo "Building $(PROJECT_NAME)..."
	go build -o bin/$(PROJECT_NAME) ./cmd/${PROJECT_NAME}/main.go

.PHONY: run
## Run the Go application
run:
	@echo "Running $(PROJECT_NAME)..."
	./bin/$(PROJECT_NAME)

.PHONY: all
## Build and the Go application
all: build run

.PHONY: deps
## Install/update Go module dependencies
deps:
	@echo "Installing/updating Go module dependencies..."
	go get ./...
	go mod tidy
	go mod download
	go mod vendor

.PHONY: test
## Run Go tests
test:
	@echo "Running tests..."
	go test ./...

.PHONY: redis-run
## Run redis locally
redis-run:
	docker run --name ${PROJECT_NAME}-redis -p ${REDIS_PORT}:6379 -d redis:7

.PHONY: redis-stop
## Run redis locally
redis-rm:
	docker rm -f ${PROJECT_NAME}-redis