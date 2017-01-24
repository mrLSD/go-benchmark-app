package main

import (
	"fmt"
	"os"
	"testing"
)

// TestLoadConfig - test load config and exactly one element
func TestLoadConfig(t *testing.T) {
	var config *Config
	config = LoadConfig()
	if len(config.Title) <= 1 {
		t.Fatal("Wrong Title")
	}
}

// TestLoadConfigWhenFileNotFound - test Load config
// with not existed config file
func TestLoadConfigWhenFileNotFound(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
				fmt.Println(err)
			}
			err = os.Rename("config/_main.toml", "config/main.toml")
			if err != nil {
				t.Fatal("Wrong Restoring file")
			}
		}
	}()

	err := os.Rename("config/main.toml", "config/_main.toml")
	if err != nil {
		t.Fatal("Wrong Renaming file")
	}
	_ = LoadConfig()
}
