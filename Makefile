# Copyright (c) 2026 SanDevil23
# SPDX-License-Identifier: Apache-2.0

APP_NAME=valentine-bot
CMD_PATH=./cmd/server
BIN_DIR=bin

.phony: all build run test clean docker linux mac windows

all: build

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
	@rm -rf $(BIN_DIR)