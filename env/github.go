package env

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/raduq/proj/config"
	"golang.org/x/oauth2"
)

// Connection to a github environment
type Connection struct {
	Client  *github.Client
	Context context.Context
}

// Connect to the github environment
func Connect() Connection {
	cfg := config.MustGet()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.GithubAccessToken})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return Connection{client, ctx}
}
