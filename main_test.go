package main

import (
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

func TestMain(t *testing.T) {
	LogFatal = func(v ...interface{}) {
		fmt.Println(v...)
	}
	main()

	// Test FAILED config
	cfg.ConfigFile = "config/_main.toml"
	main()

	// Return truly config - for next tests
	cfg.ConfigFile = "config/main.toml"
}
