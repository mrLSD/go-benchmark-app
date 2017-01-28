package main

import (
	"fmt"
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

// TestRunBanchmarks - with basic cinfig
func TestRunBanchmarls(t *testing.T) {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	// Simple check for runCommand
	_, _ = runCommand("/bin/bash")

	RunCommand = runCommandSuccess
	config.WaitToRun = 0
	config.Delay = 0

	err = RunBanchmarks(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}
}

// TestRunBanchmarksWithWrongAppPath - test with basic config
// and wrong App Path
func TestRunBanchmarksWithWrongAppPath(t *testing.T) {
	config, err := LoadConfig(CONFIG_FILE)
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
	err = RunBanchmarks(config)
	if err == nil {
		t.Fatal("Unexpected exec start result")
	}
}

// TestRunBanchmarksWithWrongParams - with basic config
// but some wrong params
func TestRunBanchmarksWithWrongParams(t *testing.T) {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	// All parametres correct
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
	config.App = []AppConfig{{
		Title: "Test Bash",
		Path:  "/bin/bash",
	}}

	// Success benchmarks
	RunCommand = runCommandSuccess
	err = RunBanchmarks(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}

	// Failed benchmarks
	RunCommand = runCommandFailed
	err = RunBanchmarks(config)
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
	err = RunBanchmarks(&abConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for abConfig")
		}
	}

	// Wrong WRK Connections parameter
	wrkConfig := *config
	wrkConfig.Wrk.Connections = 0
	err = RunBanchmarks(&wrkConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for wrkConfig")
		}
	}

	// Wrong Siege Concurrent parameter
	siegeConfig := *config
	siegeConfig.Siege.Concurrent = 0
	err = RunBanchmarks(&siegeConfig)
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

	err = RunBanchmarks(config)
	if err == nil {
		t.Fatal("Unexpected exec for KillProcess")
	}
}
