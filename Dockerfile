# -----------------------------
# Stage 1 — Build the binary
# -----------------------------
FROM golang:1.22-alpine AS builder

# Install git (needed for go modules sometimes)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy project source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o valentine-bot ./cmd/server

# -----------------------------
# Stage 2 — Run the binary
# -----------------------------
FROM alpine:latest

WORKDIR /app

# Add CA certificates (for HTTPS calls to GitHub / OpenAI)
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/github-bot .

# Expose webhook port
EXPOSE 8080

# Run the bot
CMD ["./github-bot"]