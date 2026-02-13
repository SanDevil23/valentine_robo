package auth

import (
	"github.com/google/go-github/v62/github"
)

func NewGithubClient(installationID int64) *github.Client {
	// 1. Create JWT using private key
	// 2. Exchange JWT for installation token
	// 3. Return github.NewClient(httpClient)

	return nil
}