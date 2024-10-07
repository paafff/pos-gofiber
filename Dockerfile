# Stage 1: Build the Go application
FROM golang:alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and change to the app directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build cmd/web/main.go

# Stage 2: Run the Go application
FROM alpine:latest

# Install make and other dependencies
RUN apk add --no-cache make bash

# Create and change to the app directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the config file
# COPY config.json .
COPY .env .

# Copy the config files
COPY config.production.json .

# Copy the wait-for-it.sh script
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Expose the port the app runs on
EXPOSE 3000

# Run the application
CMD ["./wait-for-it.sh", "postgres:5432","./main"]