package tools

import (
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

// TestWrkBenchCommand - test WRK command generator
func TestWrkBenchCommand(t *testing.T) {
	var tool WrkTool

	config := &cfg.Config{}
	config.Wrk.Connections = 1
	config.Wrk.Duration = 1
	config.Wrk.Threads = 1
	tool = WrkTool{&config.Wrk}
	_, err := tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Wrk.Connections = 0
	config.Wrk.Duration = 1
	config.Wrk.Threads = 1
	tool = WrkTool{&config.Wrk}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Connections")
	}

	config.Wrk.Connections = 1
	config.Wrk.Duration = 0
	config.Wrk.Threads = 1
	tool = WrkTool{&config.Wrk}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Duration")
	}

	config.Wrk.Connections = 1
	config.Wrk.Duration = 1
	config.Wrk.Threads = 0
	tool = WrkTool{&config.Wrk}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Threads")
	}
}

// TestWrkCommonResults - text common results interface
func TestWrkCommonResults(t *testing.T) {
	var tool WrkTool

	config := &cfg.Config{}
	config.Wrk.Connections = 1
	config.Wrk.Duration = 1
	config.Wrk.Threads = 1
	tool = WrkTool{&config.Wrk}
	result, err := tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}
	_ = result.Command()
	_ = result.Params()
	result.Analize()
}
