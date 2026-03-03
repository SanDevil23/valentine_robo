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

// Handle processes incoming webhook events and logs the payload.
func Handle(w http.ResponseWriter, r *http.Request) {
	payload, _ := io.ReadAll(r.Body)
	eventType := github.WebHookType(r);
	event, err := github.ParseWebHook(eventType, payload)
	if err != nil {
		log.Fatal("Error parsing webhook:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch e := event.(type){
	case *github.IssueCommentEvent:
		log.Printf("Received IssueCommentEvent: Action=%s, Comment=%s", e.GetAction(), e.Comment.GetBody())
		if e.GetAction() == "created" {
			log.Printf("New comment created: %s", e.Comment.GetBody())
			comment := e.Comment.GetBody();
			if len(comment) >= 10 && comment[:9] == "/generate" {
				requirement := comment[10:]
				go commands.HandleGenerate(
					context.Background(),
					e,
					requirement,
				)
			}
		}
	}

	w.WriteHeader(http.StatusCreated);
}