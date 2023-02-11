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
	"github.com/spf13/cobra"

	"github.com/bagmeg/ework/internal/internal_github"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "A github related actions.",
	Long:  `github 관련 작업 진행`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}

// ---------------------- list -------------------------------

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list github repositories",
	Run:   internal_github.ListRun,
}

// ---------------------- repository -------------------------------

var cloneCmd = &cobra.Command{
	Use:   "clone repository_name",
	Short: "clone given repository to current directory",
	Run:   internal_github.CloneRun,
}

// ---------------------- init -------------------------------

func init() {
	rootCmd.AddCommand(githubCmd)

	// -----------------------------------------------------

	githubCmd.AddCommand(listCmd)
	githubCmd.AddCommand(cloneCmd)

	// -----------------------------------------------------

	listCmd.Flags().BoolP("organization", "o", false, "list organizations")
	listCmd.Flags().BoolP("repository", "r", false, "list repositories`")

	// -----------------------------------------------------
}
