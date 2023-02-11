/*
Copyright © 2023 bagmeg
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bagmeg/ework/internal/config"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ework",
	Short: "exem work command line tool",
	Long:  `A worker for the exem`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello ework")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", config.DefaultConfigPath, "config file (default is $HOME/.ework/config.yaml)")
}
