// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"io"
	"log"
	"net/http"
)

// Handle processes incoming webhook events and logs the payload.
func Handle(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	log.Println("Webhook event received:", string(body))
	w.WriteHeader(http.StatusOK)
}