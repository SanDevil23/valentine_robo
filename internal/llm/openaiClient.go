package llm

import (
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