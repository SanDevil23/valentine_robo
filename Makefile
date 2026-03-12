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

.phony: all build run test clean docker linux mac windows

all: clean build run

build:
	@echo "Building $(APP_NAME)....."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)
	@echo "Build successful for $(APP_NAME)....."


run:
	@ngrok http 8080
	@go run $(CMD_PATH)
	echo "Running $(APP_NAME)....."

clean:
	@echo "Cleaning up $(BIN_DIR)....."
	@rm -rf $(BIN_DIR)
	@echo "Cleaned $(BIN_DIR)....."


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

# -----------------------------
# Full Rebuild
# -----------------------------
rebuild: docker-stop docker-build docker-run
How to Use It
Build Go binary