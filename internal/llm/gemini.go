// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0

package llm

import (
	"context"
	"os"

	"google.golang.org/api/option"
	"github.com/google/generative-ai-go/genai"
)

// Client wraps the Gemini client
type GeminiClient struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// GetNewClient initializes Gemini client
func GetGeminiClient(ctx context.Context) (*GeminiClient, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")

	c, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	model := c.GenerativeModel("gemini-1.5-flash") // fast & cheap equivalent

	return &GeminiClient{
		client: c,
		model:  model,
	}, nil
}

// GenerateCode generates structured code output
func (cl *GeminiClient) GenerateCode(ctx context.Context, requirement string) (string, error) {
	prompt := `You are a senior Golang engineer. Generate a production-ready project for:` + requirement + `
Output ONLY in this JSON format:
{
  "files": [
    {
      "path": "file path here",
      "content": "code here"
    }
  ]
}`

	resp, err := cl.model.GenerateContent(
		ctx,
		genai.Text(prompt),
	)

	if err != nil {
		return "Failed to generate code", err
	}

	// Gemini responses are structured differently
	var output string
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				output += string(txt)
			}
		}
	}

	return output, nil
}