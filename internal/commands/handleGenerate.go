package commands

import (
	"context"
	"log"

	"github.com/google/go-github/v62/github"
	githubclient "github.com/sandevil23/valentin_robo/internal/github"
	githubclient_auth "github.com/sandevil23/valentin_robo/internal/github/auth"
	"github.com/sandevil23/valentin_robo/internal/llm"
)

func HandleGenerate(ctx context.Context, e *github.IssueCommentEvent, req string) {

	installationID := e.GetInstallation().GetID();
	files := make(map[string]string)

	client, err := githubclient_auth.NewGithubClient(installationID)
	if err != nil {
		log.Println("Error creating GitHub client:", err)
		return
	}
	
	llmClient := llm.NewHFClient()

	// 1. Generate code
	raw, err := llmClient.GenerateCode(ctx, req)
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
		files[f.Path] = f.Content
	}

	owner := e.GetRepo().GetOwner().GetLogin()
	repo := e.GetRepo().GetName()

	err = githubclient.CreatePR(ctx, client, owner, repo, "valentine-bot", files)
	if err != nil{
		log.Fatalln("Failed to create pull request")
		log.Fatalln(err)
	}
}
