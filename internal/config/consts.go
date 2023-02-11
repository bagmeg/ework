package config

import "os"

func init() {
	// Default path is $HOME/.ework/config.yaml
	homeDir, _ := os.UserHomeDir()
	ConfigPath = homeDir + DefaultConfigPath
}

var (
	ConfigPath string
)

const (
	DefaultConfigPath = "~/.ework/config.yaml"
)

const (
	GithubToken = "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"
)
