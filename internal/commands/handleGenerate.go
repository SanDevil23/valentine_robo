package commands

import (
	"context"
	"log"

	"github.com/sandevil23/valentin_robo/internal/llm"
)

func HandleGenerate(ctx context.Context, requirement string) {

	client := llm.GetNewClient()

	// 1. Generate code
	raw, err := client.GenerateCode(ctx, requirement)
	if err != nil {
		log.Println("LLM error:", err)
		return
	}

	// 2. Parse response
	res, err := llm.ParseResponse(raw)
	if err != nil {
		log.Println("Parse error:", err)
		return
	}

	// 3. Validate files
	err = llm.ValidateFiles(res.Files)
	if err != nil {
		log.Println("Validation error:", err)
		return
	}

	// 4. Create files in GitHub (your existing code)
	for _, f := range res.Files {
		log.Println("Creating file:", f.Path)

		// TODO: call your GitHub client here
	}
}
