# Copyright (c) 2026 SanDevil23
# SPDX-License-Identifier: Apache-2.0

# ==========================================================
# Project Configuration
# ==========================================================

APP_NAME=valentine-bot
BINARY_NAME := $(APP_NAME)
CMD_PATH=./cmd/server
BIN_DIR=bin

VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

IMAGE_NAME=valentine-bot
DOCKER_REGISTRY := docker.io/yourusername
CONTAINER_NAME=valentine-bot-container
PORT=8080

GO := go

# ==========================================================
# Help
# ==========================================================

.PHONY: help
help:
	@echo "Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make run           Run service locally"
	@echo "  make build         Build binary"
	@echo "  make clean         Remove build artifacts"
	@echo ""
	@echo "Code Quality:"
	@echo "  make fmt           Format code"
	@echo "  make lint          Run linter"
	@echo "  make test          Run tests"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build  Build Docker image"
	@echo "  make docker-run    Run container"
	@echo "  make docker-stop   Stop container"
	@echo "  make docker-push   Push image to registry"
	@echo ""
	@echo "Release:"
	@echo "  make release       Tag and push release"


.phony: all build run test clean docker linux mac windows

all: clean build run

build:
	@echo "Building $(APP_NAME)....."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)
	@echo "Build successful for $(APP_NAME)....."


run:
	echo "Running $(APP_NAME)....."
	@go run $(CMD_PATH)

clean:
	@echo "Cleaning up $(BIN_DIR)....."
	@rm -rf $(BIN_DIR)
	@echo "Cleaned $(BIN_DIR)....."


.PHONY: tunnel
tunnel:
	ngrok http 8080

# ==========================================================
# Code Quality
# ==========================================================

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	$(GO) test -v ./...

# -----------------------------
# Docker Build
# -----------------------------
docker-build:
	docker build -t $(IMAGE_NAME):latest .

# -----------------------------
# Docker Run
# -----------------------------
docker-run:
	docker run -p $(PORT):8080 \
	--name $(CONTAINER_NAME) \
	-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
	-e OPENAI_API_KEY=$(OPENAI_API_KEY) \
	$(IMAGE_NAME):latest

# -----------------------------
# Docker Stop
# -----------------------------
docker-stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

# -----------------------------
# Docker Restart
# -----------------------------
docker-restart: docker-stop docker-run

# -----------------------------
# Logs
# -----------------------------
logs:
	docker logs -f $(CONTAINER_NAME)

# ==========================================================
# Utilities
# ==========================================================

.PHONY: rebuild
rebuild: clean build

.PHONY: dev
dev: fmt build run
# -----------------------------
# Full Rebuild
# -----------------------------
rebuild: docker-stop docker-build docker-run
