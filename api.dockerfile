# Stage 1: Build stage
FROM golang:1.21-alpine AS builder

# Install required build tools
RUN apk add --no-cache gcc musl-dev git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application with more verbose output
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/server/main.go

# Stage 2: Final stage
FROM alpine:3.19

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates postgresql-client

WORKDIR /app

# Copy the built executable from builder
COPY --from=builder /app/main .

# Copy necessary configuration and migration files
COPY .env .
COPY ./migrations ./db/migrations

# Create a non-root user
RUN adduser -D appuser && \
    chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose the port the server listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
