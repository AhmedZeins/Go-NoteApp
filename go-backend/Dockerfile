# Build stage
FROM golang:1.21-alpine3.18 AS builder

# Install security scanner and build dependencies
RUN apk add --no-cache ca-certificates git

# Create a non-root user
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with security flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /go/bin/app ./cmd/api

# Final stage
FROM alpine:3.18

# Add security updates and CA certificates
RUN apk --no-cache add ca-certificates tzdata \
    && update-ca-certificates

# Import the user from builder
COPY --from=builder /etc/passwd /etc/passwd

# Copy the binary from builder
COPY --from=builder /go/bin/app /app

# Copy static files if needed
COPY --from=builder /app/static /static

# Use non-root user
USER appuser

# Set environment variables
ENV GIN_MODE=release

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/app"]
