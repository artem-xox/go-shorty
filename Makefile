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

.PHONY: reqs
## Install/update additional requirenments
reqs:
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.52.0

.PHONY: test
## Run Go tests
test:
	@echo "Running tests..."
	go test ./...

.PHONY: lint
## Run linter
lint:
	@echo "Running code lint..."
	./bin/golangci-lint run 

.PHONY: redis-run
## Run redis locally
redis-run:
	docker run --name ${PROJECT_NAME}-redis -p ${REDIS_PORT}:6379 -d redis:7

.PHONY: redis-stop
## Remove redis locally
redis-rm:
	docker rm -f ${PROJECT_NAME}-redis

.PHONY: ping
## Ping local service
ping:
	curl -I 0.0.0.0:8081/ping

.PHONY: test-pipeline
##
test-pipeline:
	curl -X POST -H 'Content-Type: application/json' -d '{"url": "https://www.facebook.com"}' 0.0.0.0:8081/set
	curl -X POST -H 'Content-Type: application/json' -d '{"url": "https://www.google.com"}' 0.0.0.0:8081/set
	curl -X POST -H 'Content-Type: application/json' -d '{"url": "https://www.github.com/artem-xox"}' 0.0.0.0:8081/set

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down