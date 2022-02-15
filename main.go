package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v42/github"
	"golang.org/x/oauth2"
)

func githubClient(ctx context.Context, token string) *github.Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return github.NewClient(oauth2.NewClient(ctx, tokenSource))
}

func main() {
	ghToken := os.Getenv("PLUGIN_GITHUB_TOKEN")
	commentBody := os.Getenv("PLUGIN_COMMENT")

	issueNumberStr, exists := os.LookupEnv("PLUGIN_ISSUE_NUMBER")

	if !exists {
		issueNumberStr = os.Getenv("DRONE_PULL_REQUEST")
	}

	issueNumber, err := strconv.Atoi(issueNumberStr)
	if err != nil {
		fmt.Printf("Error converting issue number [%s] to int\n", issueNumberStr)
		os.Exit(1)
	}
	namespace := os.Getenv("DRONE_REPO_NAMESPACE")
	repo := os.Getenv("DRONE_REPO_NAME")

	ctx := context.Background()
	client := githubClient(ctx, ghToken)

	issueComment := &github.IssueComment{
		Body: &commentBody,
	}
	log.Printf("Commenting on issue [%d] in repo [%s/%s]\n", issueNumber, namespace, repo)
	_, resp, err := client.Issues.CreateComment(
		context.Background(),
		namespace,
		repo,
		issueNumber,
		issueComment,
	)
	if err != nil {
		fmt.Printf("Error creating comment in Github [%s]\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != 201 {
		fmt.Printf("Error creating comment in Github [%s]\n", resp.Status)
		os.Exit(1)
	}
}
