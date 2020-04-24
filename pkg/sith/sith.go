package sith

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewGithubClient creates a new Github Client
func NewGithubClient(ctx context.Context) (*github.Client, error) {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		err := fmt.Errorf("No Access token provided, export GITHUB_TOKEN=%s", "'valid token")
		return nil, err
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client, nil
}
