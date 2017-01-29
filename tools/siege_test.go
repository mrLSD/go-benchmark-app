package tools

import (
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

// TestSiegeBenchCommand - test Siege command generator
func TestSiegeBenchCommand(t *testing.T) {
	var tool SiegeTool

	config := &cfg.Config{}
	config.Siege.Concurrent = 1
	config.Siege.Time = 1
	tool = SiegeTool{&config.Siege}
	_, err := tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Siege.Concurrent = 0
	config.Siege.Time = 1
	tool = SiegeTool{&config.Siege}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Concurrent")
	}

	config.Siege.Concurrent = 1
	config.Siege.Time = 0
	tool = SiegeTool{&config.Siege}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Time")
	}
}

// TestSiegeCommonResults - text common results interface
func TestSiegeCommonResults(t *testing.T) {
	var tool SiegeTool

	config := &cfg.Config{}
	config.Siege.Concurrent = 1
	config.Siege.Time = 1
	tool = SiegeTool{&config.Siege}
	result, err := tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}
	_ = result.Command()
	_ = result.Params()
	result.Analize()
}
