// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/google/go-github/v62/github"
	"github.com/sandevil23/valentin_robo/internal/commands"
)

// Handle is an HTTP handler that processes incoming GitHub webhook events.
//
// It reads the request payload, determines the GitHub event type,
// parses the webhook into a strongly-typed GitHub event struct,
// and conditionally dispatches supported events to internal command handlers.
//
// Currently supported events:
//   - IssueCommentEvent: triggers command processing when a new comment
//     begins with a recognized slash command (e.g., "/generate").
//
// The handler returns:
//   - HTTP 400 if webhook parsing fails
//   - HTTP 201 after successful processing
func Handle(w http.ResponseWriter, r *http.Request) {
	// Read the raw request body sent by GitHub.
	// Note: In production systems, error handling on ReadAll should not be ignored.
	payload, _ := io.ReadAll(r.Body)

	// Determine the GitHub webhook event type from headers.
	eventType := github.WebHookType(r)

	// Parse the webhook payload into a concrete event structure.
	event, err := github.ParseWebHook(eventType, payload)
	if err != nil {
		// Fatal is used here to ensure visibility of malformed payloads.
		// In high-availability systems, consider structured logging instead.
		log.Fatal("Error parsing webhook:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Perform type-switch to handle only supported GitHub event types.
	switch e := event.(type) {

	// Handle issue comment events.
	case *github.IssueCommentEvent:

		// Log metadata for observability and debugging.
		log.Printf(
			"Received IssueCommentEvent: Action=%s, Comment=%s",
			e.GetAction(),
			e.Comment.GetBody(),
		)

		// Process only newly created comments.
		if e.GetAction() == "created" {

			log.Printf("New comment created: %s", e.Comment.GetBody())

			comment := e.Comment.GetBody()

			// Check if the comment begins with the supported slash command.
			// This acts as a lightweight command router.
			if len(comment) >= 10 && comment[:9] == "/generate" {

				// Extract the command argument (requirement) following the prefix.
				requirement := comment[10:]

				// Dispatch command processing asynchronously.
				// A goroutine is used to avoid blocking the webhook response cycle.
				// This ensures GitHub receives a timely HTTP response.
				go commands.HandleGenerate(
					context.Background(),
					e,
					requirement,
				)
			}
		}
	}

	// Return 201 Created to acknowledge successful webhook handling.
	// GitHub expects a 2xx response to consider delivery successful.
	w.WriteHeader(http.StatusCreated)
}