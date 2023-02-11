package config

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T) {
	c := New()

	f, err := os.ReadFile("../../test/config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	yaml.Unmarshal(f, &c)

	fmt.Println(c)
}
