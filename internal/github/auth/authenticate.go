package auth
// Copyright (c) 2026 SanDevil23
// SPDX-License-Identifier: Apache-2.0
import (
	"context"
	"encoding/pem"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"
)

func NewGithubClient(installationID int64) (*github.Client, error) {
	// 1. Create JWT using private key
	// 2. Exchange JWT for installation token
	// 3. Return github.NewClient(httpClient)

	ctx := context.Background();
	appIdStr := os.Getenv("GITHUB_APP_ID")
	if (appIdStr == "") {
		return nil, errors.New("GITHUB_APP_ID environment variable is not set")
	}

	// parse appIdStr to int64
	appId, err := strconv.ParseInt(appIdStr, 10, 64)
	if err != nil {
		return nil, errors.New("GITHUB_APP_ID environment variable is not a valid integer")
	}

	privateKeyPEM := os.Getenv("PRIVATE_KEY")
	if (privateKeyPEM == "") {
		return nil, errors.New("PRIVATE_KEY environment variable is not set")
	}

	block, _ := pem.Decode(([]byte(privateKeyPEM)))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))	
	if privateKey == nil {
		return nil, errors.New("failed to parse RSA private key")
	}	
	
	// create JWT Token
	now := time.Now()
	claims := jwt.RegisteredClaims{
		IssuedAt: jwt.NewNumericDate(now.Add(-1 * time.Minute)),
		ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Minute)),
		Issuer: strconv.FormatInt(appId, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwtString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	// create a temporary HTTP client with the JWT token
	jwtTransport := &http.Transport{}
	jwtClient := &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set("Authorization", "Bearer "+jwtString)
			req.Header.Set("Accept", "application/vnd.github+json")
			return jwtTransport.RoundTrip(req)
		}),
	}

	appClient := github.NewClient(jwtClient)

	// exchange JWT for installation token
	installationToken, _ , err := appClient.Apps.CreateInstallationToken(
		ctx,
		installationID,
		&github.InstallationTokenOptions{},
	)
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: installationToken.GetToken()},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc), nil
}


// Helper type to inject headers
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}