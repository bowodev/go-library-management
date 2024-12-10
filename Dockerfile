# Step 1: Build stage
FROM golang:1.22.4-alpine3.20 AS builder

# Install necessary dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd

# Step 2: Runtime stage
FROM alpine:latest

# Install certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

COPY config.yaml .

# Expose application port
EXPOSE 8082

# Command to run the application
CMD ["./main"]
