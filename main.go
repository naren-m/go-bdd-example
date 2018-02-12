package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "1161bc4a9c064328d5693421207ae717c129ffec"},
	)

	tc := oauth2.NewClient(ctx, ts)
	// get go-github client
	client := github.NewClient(tc)

	repo, _, err := client.Repositories.Get(ctx, "Golang-Coach", "Lessons")
	//	repo, _, err := client.Repositories.Get(ctx, "naren-m", "dotfiles")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)

}
