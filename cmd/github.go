/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

// NOTE:
// 1. github에서 지정한 repository를 clone한다.
// 2. github에서 내가 속한 organization의 repository를 clone한다.
// 3. organization의 repository를 clone할때는 organization의 이름을 입력받아야 한다.
// 4. organization의 repository를 clone할때는 특정 기간동안 action이 없는 repository는 clone하지 않는다.

import (
	"context"
	"fmt"

	"github.com/bagmeg/ework/cmd/config"
	"github.com/google/go-github/v50/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	list string
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "A github related actions.",
	Long:  `github 관련 작업 진행`,
	Run:   run,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list github repositories",
	Long:  "list github repositories",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: config.GithubToken},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		// list all repositories for the authenticated user
		repos, _, err := client.Repositories.List(ctx, "", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, repo := range repos {
			fmt.Println(*repo.Name)
		}
	},
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("github called")

}

func init() {
	rootCmd.AddCommand(githubCmd)

	githubCmd.AddCommand(listCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
