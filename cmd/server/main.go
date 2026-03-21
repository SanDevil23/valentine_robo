// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"log"
	"net/http"

	// "github.com/sandevil23/valentin_robo/internal/commands"
	"github.com/joho/godotenv"
	"github.com/sandevil23/valentin_robo/internal/controllers"
	"github.com/sandevil23/valentin_robo/internal/webhook"
)

func main() {
	err := godotenv.Load()
	if err!=nil{
		log.Fatalln("Failed to load environment variables")
	}

	// routes
	http.HandleFunc("/webhook", webhook.Handle)
	http.HandleFunc("/", controllers.HomeHandler)

	log.Println("Valentine Robo server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// commands.IsGenerateCommand("/generate list APIs");
}