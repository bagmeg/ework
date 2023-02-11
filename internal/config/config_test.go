package config

import (
	"fmt"
	"testing"
)

const (
	testDir = "../../test"
)

func TestConfig(t *testing.T) {
	c := New()

	if err := Load(c, []string{testDir + "/conig.yaml"}); err != nil {
		t.Fatal(err)
	}

	fmt.Println(*c)
}
