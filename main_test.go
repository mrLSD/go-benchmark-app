package main

import (
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

// runBenchmarksSuccess - alias for success Run Benchmarks
var runBenchmarksSuccess = func(config *cfg.Config) error {
	println("	=> runBenchmarksSuccess")
	return nil
}

// runBenchmarksFailed - alias for failed Run Benchmarks
var runBenchmarksFailed = func(config *cfg.Config) error {
	println("	=> runBenchmarksFailed")
	return fmt.Errorf("test %s", "test")
}

func TestMain(t *testing.T) {
	RunBenchmarks = runBenchmarksSuccess
	LogFatal = func(v ...interface{}) {
		fmt.Println(v...)
	}
	main()

	// Test FAILED config
	cfg.ConfigFile = "config/_main.toml"
	main()

	// Return truly config - for next tests
	cfg.ConfigFile = "config/main.toml"
	// Test Benchmarks Failed
	RunBenchmarks = runBenchmarksFailed
	main()
}
