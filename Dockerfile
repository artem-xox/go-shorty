# Use the Golang 1.18.1 image as the base image
FROM golang:1.18.1 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o shorty cmd/shorty/main.go

# Use a smaller, official image for the final build
FROM ubuntu:latest

# Set the working directory to /app
WORKDIR /app

# Copy the built Go application from the builder image to the final image
COPY --from=builder /app/shorty /go/bin/shorty

EXPOSE 8081

# Set an entrypoint to run the Go application
ENTRYPOINT ["/go/bin/shorty"]

