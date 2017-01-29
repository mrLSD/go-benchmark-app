package main

import (
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	LogFatal = func(v ...interface{}) {
		fmt.Println(v...)
	}
	main()

	// Test FAILED config
	if err := os.Rename(cfg.CONFIG_FILE, "config/_main.toml"); err != nil {
		t.Fatal(err)
	}
	defer os.Rename("config/_main.toml", cfg.CONFIG_FILE)
	main()
}
