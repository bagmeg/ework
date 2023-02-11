package internal_github

import (
	"context"
	"fmt"

	"github.com/bagmeg/ework/internal/config"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	targetFuncs []func(*cobra.Command, *github.Client, context.Context)
)

func init() {
	targetFuncs = append(targetFuncs, listRepositories)
	targetFuncs = append(targetFuncs, listOrganizations)
}

func ListRun(cmd *cobra.Command, args []string) {
	cfg := config.New()
	cf, err := cmd.Flags().GetString("config")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := config.Load(cfg, []string{cf}); err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	for _, f := range targetFuncs {
		f(cmd, client, ctx)
	}
}

func listRepositories(cmd *cobra.Command, client *github.Client, ctx context.Context) {
	listRepos, err := cmd.Flags().GetBool("repository")
	if err != nil {
		fmt.Println(err)
		return
	}

	if listRepos == false {
		return
	}

	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}

func listOrganizations(cmd *cobra.Command, client *github.Client, ctx context.Context) {
	listOrgs, err := cmd.Flags().GetBool("organization")
	if err != nil {
		fmt.Println(err)
		return
	}

	if listOrgs == false {
		return
	}

	orgs, _, err := client.Organizations.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, org := range orgs {
		fmt.Println(*org.Name)
	}
}
