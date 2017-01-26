package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	LogFatal = func(v ...interface{}) {
		fmt.Println(v...)
	}
	main()

	// Test FAILED config
	if err := os.Rename(CONFIG_FILE, "config/_main.toml"); err != nil {
		t.Fatal(err)
	}
	defer os.Rename("config/_main.toml", CONFIG_FILE)
	main()
}
