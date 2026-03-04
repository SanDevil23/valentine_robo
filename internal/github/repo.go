package githubclient

import (
	"context"
	"log"

	"github.com/google/go-github/v62/github"
)

func CreatePR(
	ctx context.Context,
	client *github.Client,
	owner, repo, branch string,
	files map[string]string,
) error {

	baseRef,_, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/main")
	if err != nil {
		log.Println("Error getting base ref:", err)
		return err;
	}

	newRef := &github.Reference{
		Ref: github.String("refs/heads/" + branch),
		Object: &github.GitObject{
			SHA: baseRef.Object.SHA,
		},
	}

	client.Git.CreateRef(ctx, owner, repo, newRef)
	return err
}
