// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"
	"github.com/sandevil23/valentin_robo/internal/webhook"
)

func main() {
	http.HandleFunc("/webhook", webhook.Handle)
	log.Println("Valentin Robo server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}