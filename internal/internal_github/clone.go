package internal_github

import (
	"context"
	"fmt"

	"github.com/bagmeg/ework/internal/config"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func init() {

}

func CloneRun(cmd *cobra.Command, args []string) {
	if err := checkValide(args); err != nil {
		fmt.Println(err)
		return
	}

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

	repo, _, err := client.Repositories.Get(ctx, cfg.User.Name, args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	url := repo.GetCloneURL()
	fmt.Printf("Clone url: %s\n", url)

	_, err = git.PlainClone("/home/ubuntu/ework/testDir/hecto", false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: cfg.User.Name,
			Password: cfg.Token,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func checkValide(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Repository name is required")
	}
	return nil
}
