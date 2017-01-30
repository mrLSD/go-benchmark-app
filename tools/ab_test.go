package tools

import (
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

// TestAbBenchCommand - test AB command generator
func TestAbBenchCommand(t *testing.T) {
	var tool AbTool

	config := &cfg.Config{}
	config.Ab.Keepalive = false
	config.Ab.Concurency = 1
	config.Ab.Requests = 1

	tool = AbTool{&config.Ab}
	_, err := tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Ab.Keepalive = true
	tool = AbTool{&config.Ab}
	_, err = tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Ab.Concurency = 2
	tool = AbTool{&config.Ab}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Requests < Concurency")
	}

	config.Ab.Concurency = 0
	tool = AbTool{&config.Ab}
	_, err = tool.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Requests = 0")
	}
}

// TestAbCommonResults - text common results interface
func TestAbCommonResults(t *testing.T) {
	var tool AbTool

	config := &cfg.Config{}
	config.Ab.Keepalive = false
	config.Ab.Concurency = 1
	config.Ab.Requests = 1
	tool = AbTool{&config.Ab}
	result, err := tool.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}
	_ = result.Command()
	_ = result.Params()
	data := []byte("")
	result.Analyze(data)
}
