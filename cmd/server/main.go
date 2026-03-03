// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/sandevil23/valentin_robo/internal/webhook"
	// "github.com/sandevil23/valentin_robo/internal/commands"
)

func main() {
	_ = godotenv.Load();

	http.HandleFunc("/webhook", webhook.Handle)
	
	log.Println("Valentine Robo server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}