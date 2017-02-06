package config

import (
	"testing"
)

// TestLoadConfig - test load config and exactly one element
func TestLoadConfig(t *testing.T) {
	var config *Config
	config, err := LoadConfig("main.toml", config)
	if err != nil {
		t.Fatal(err)
	}
	if len(config.Title) <= 1 {
		t.Fatal("Wrong Title")
	}
}

// TestLoadConfigWhenFileNotFound - test Load config
// with not existed config file
func TestLoadConfigWhenFileNotFound(t *testing.T) {
	var err error
	var config *Config
	_, err = LoadConfig("main.toml", config)
	if err != nil {
		t.Fatal(err)
	}

	_, err = LoadConfig("config/_main.toml", config)
	if err == nil {
		t.Fatal("Anexpected for wrong config")
	}
}
