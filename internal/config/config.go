package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	DefaultToken = ""
)

var (
	DefaultConfig = Config{
		Token: DefaultToken,
		User:  DefaultUser,
	}
)

type Config struct {
	Token string `yaml:"GithubToken"`
	User  User   `yaml:"User,omitempty"`
}

func New() *Config {
	return &DefaultConfig
}

func Load(c *Config, paths []string) error {
	for _, p := range paths {
		if err := checkValid(p); err != nil {
			return fmt.Errorf("invalid config path: %w", err)
		}
	}

	for _, p := range paths {
		file, err := os.Open(p)
		defer file.Close()
		if err != nil {
			return err
		}

		if err := yaml.NewDecoder(file).Decode(c); err != nil {
			return err
		}
	}
	return nil
}

func checkValid(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("config path is empty")
	}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("config path is a directory")
	}

	return nil
}
