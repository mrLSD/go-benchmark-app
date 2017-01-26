package main

import (
	"testing"
)

func TestRunBanchmars(t *testing.T) {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	err = RunBanchmars(config)
	if err != nil {
		t.Fatal(err)
	}

	if len(config.App) > 0 {
		config.App[0].Path = "test/test"
	} else {
		t.Fatal("You should have at least one App")
	}
	err = RunBanchmars(config)
	if err == nil {
		t.Fatal("Unexpected exec start result")
	}
}
