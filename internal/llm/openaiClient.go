package llm

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

func GetNewClient() *Client  {
	return &Client{
		client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
	}
}

func (cl *Client) GenerateCode(ctx context.Context, requirement string) (string, error) {
	prompt := `You are a senior Golang engineer. Generate a production-ready project for:` + requirement + `Output ONLY in this JSON format:
				{
				"files": [
				{
					"path": "file path here",
					"content": "code here"
				}
				]
				}
	`

	resp, err := cl.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini, // fast & cheap
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "You are an expert backend engineer",
				},
				{
					Role:    "user",
					Content: prompt,
				},
			},
			Temperature: 0.2,
		},
	)

	if err != nil {
		return "Failed to generate code", err
	}

	return resp.Choices[0].Message.Content, nil

}