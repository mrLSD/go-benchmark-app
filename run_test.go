package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// TestRunBanchmars - with basic cinfig
func TestRunBanchmarls(t *testing.T) {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		t.Fatal(err)
	}

	err = RunBanchmars(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}
}

// TestRunBanchmarsWithWrongAppPath - test with basic config
// and wrong App Path
func TestRunBanchmarsWithWrongAppPath(t *testing.T) {
	config, err := LoadConfig(CONFIG_FILE)
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

// TestRunBanchmarsWithWrongParams - with basic config
// but some wrong params
func TestRunBanchmarsWithWrongParams(t *testing.T) {
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

	// Re-init app
	config.App = []AppConfig{AppConfig{
		Title: "Tset Bash",
		Path:  "/bin/bash",
	}}

	err = RunBanchmars(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}

	// Wrong AB Concurency parameter
	abConfig := *config
	abConfig.Ab.Concurency = 0
	err = RunBanchmars(&abConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for abConfig")
		}
	}

	// Wrong WRK Connections parameter
	wrkConfig := *config
	wrkConfig.Wrk.Connections = 0
	err = RunBanchmars(&wrkConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for wrkConfig")
		}
	}

	// Wrong Siege Concurrent parameter
	siegeConfig := *config
	siegeConfig.Siege.Concurrent = 0
	err = RunBanchmars(&siegeConfig)
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

	err = RunBanchmars(config)
	if err == nil {
		t.Fatal("Unexpected exec for KillProcess")
	}
}
