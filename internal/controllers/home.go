// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0
package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Message string `json:"message"`
	Status string `json:"status"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	response := Response {
		Message: "Welcome to Valentine Robo",
		Status: http.StatusText(200),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}