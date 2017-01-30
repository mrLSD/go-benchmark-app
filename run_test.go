package main

import (
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"os"
	"os/exec"
	"testing"
)

// Alias for success runned command
var runCommandSuccess = func(c string, args ...string) ([]byte, error) {
	println("	=> runCommandSuccess")
	return []byte("test"), nil
}

// Alias for failed runned command
var runCommandFailed = func(c string, args ...string) ([]byte, error) {
	println("	=> runCommandFailed")
	return []byte("test"), fmt.Errorf("test %s", "test")
}

// TestRunBenchmarks - with basic cinfig
func TestRunBenchmarks(t *testing.T) {
	config, err := cfg.LoadConfig(cfg.CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	// Simple check for runCommand
	_, _ = runCommand("/bin/bash")

	RunCommand = runCommandSuccess
	config.WaitToRun = 0
	config.Delay = 0

	err = RunBenchmarks(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}
}

// TestRunBenchmarksWithWrongAppPath - test with basic config
// and wrong App Path
func TestRunBenchmarksWithWrongAppPath(t *testing.T) {
	config, err := cfg.LoadConfig(cfg.CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	config.WaitToRun = 0
	config.Delay = 0

	if len(config.App) > 0 {
		config.App[0].Path = "test/test"
	} else {
		t.Fatal("You should have at least one App")
	}
	err = RunBenchmarks(config)
	if err == nil {
		t.Fatal("Unexpected exec start result")
	}
}

// TestRunBenchmarksWithWrongParams - with basic config
// but some wrong params
func TestRunBenchmarksWithWrongParams(t *testing.T) {
	config, err := cfg.LoadConfig(cfg.CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	// All parameters correct
	config.Ab.Requests = 1
	config.Ab.Concurency = 1
	config.Wrk.Threads = 1
	config.Wrk.Connections = 1
	config.Wrk.Duration = 1
	config.Siege.Concurrent = 1
	config.Siege.Time = 1

	config.WaitToRun = 0
	config.Delay = 0

	// Re-init app
	config.App = []cfg.AppConfig{{
		Title: "Test Bash",
		Path:  "/bin/bash",
	}}

	// Success benchmarks
	RunCommand = runCommandSuccess
	err = RunBenchmarks(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}

	// Failed benchmarks
	RunCommand = runCommandFailed
	err = RunBenchmarks(config)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for runCommandFailed")
		}
	}

	// Return to Success Run command
	RunCommand = runCommandSuccess

	// Wrong AB Concurency parameter
	abConfig := *config
	abConfig.Ab.Concurency = 0
	err = RunBenchmarks(&abConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for abConfig")
		}
	}

	// Wrong WRK Connections parameter
	wrkConfig := *config
	wrkConfig.Wrk.Connections = 0
	err = RunBenchmarks(&wrkConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for wrkConfig")
		}
	}

	// Wrong Siege Concurrent parameter
	siegeConfig := *config
	siegeConfig.Siege.Concurrent = 0
	err = RunBenchmarks(&siegeConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for siegeConfig")
		}
	}

	// Simulate Wrong Kill Process
	KillProcess = func(cmd *exec.Cmd) error {
		return fmt.Errorf("test %s", "test")
	}

	err = RunBenchmarks(config)
	if err == nil {
		t.Fatal("Unexpected exec for KillProcess")
	}
}
