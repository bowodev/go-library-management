# Base image
FROM golang:1.22.4-alpine3.20

# Install necessary dependencies
RUN apk add --no-cache git curl

# Install Air for hot reload
RUN curl -fLo /usr/bin/air https://github.com/cosmtrek/air/releases/download/v1.44.0/air && chmod +x /usr/bin/air

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Expose application port
EXPOSE 8082

# Start the application with Air
CMD ["air", "-c", ".air.toml"]
