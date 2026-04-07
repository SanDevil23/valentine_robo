package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// HFClient represents the Hugging Face LLM client
type HFClient struct {
	apiKey string
	model  string
	client *http.Client
}

// NewHFClient initializes a Hugging Face client using 7B Llama 3
func NewHFClient() *HFClient {
	return &HFClient{
		apiKey: os.Getenv("HF_API_KEY"),
		model:  "meta-llama/Llama-3.1-8B-Instruct:novita",
		client: &http.Client{
			Timeout: 120 * time.Second, // allow longer model load
		},
	}
}

// hfRequest is the request body for Hugging Face inference API
type hfRequest struct {
	Inputs  string                 `json:"inputs"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// hfResponse is the response format for HF models
type hfResponse []struct {
	GeneratedText string `json:"generated_text"`
}

// GenerateCode generates production-ready Go project JSON
func (c *HFClient) GenerateCode(ctx context.Context, requirement string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("HF_API_KEY not set")
	}

	prompt := fmt.Sprintf(`<s>[INST] You are a senior Golang engineer.

Generate a production-ready Go project for:

"""
%s
"""

Output ONLY valid JSON in this format:
{
  "files": [
    {
      "path": "file path here",
      "content": "code here"
    }
  ]
}
[/INST]`, requirement)

	reqBody := hfRequest{
		Inputs: prompt,
		Options: map[string]interface{}{
			"wait_for_model": true,
			"max_new_tokens": 1024,
			"temperature":    0.2,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	// Use new HF router endpoint
	url := fmt.Sprintf("https://router.huggingface.co/hf-inference/models/%s", c.model)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-use-cache", "false") // optional, improves stability

	// Retry loop for transient errors
	for i := 0; i < 3; i++ {
		resp, err := c.client.Do(req)
		if err != nil {
			time.Sleep(time.Duration(i+1) * 2 * time.Second)
			continue
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			// Log body for debugging
			return "", fmt.Errorf("HF API error (status %d): %s", resp.StatusCode, string(bodyBytes))
		}

		var result hfResponse
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			return "", fmt.Errorf("failed to decode response: %v | body: %s", err, string(bodyBytes))
		}

		if len(result) == 0 {
			return "", errors.New("empty response from model")
		}

		return result[0].GeneratedText, nil
	}

	return "", errors.New("failed to get response from HF after 3 retries")
}