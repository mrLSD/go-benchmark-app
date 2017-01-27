package main

import (
	"testing"
)

// TestAbBenchCommand - test AB command generator
func TestAbBenchCommand(t *testing.T) {
	config := &Config{}
	config.Ab.Keepalive = false
	config.Ab.Concurency = 1
	config.Ab.Requests = 1
	_, _, err := config.Ab.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Ab.Keepalive = true
	_, _, err = config.Ab.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Ab.Concurency = 2
	_, _, err = config.Ab.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Requests < Concurency")
	}

	config.Ab.Concurency = 0
	_, _, err = config.Ab.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Requests = 0")
	}
}

// TestWrkBenchCommand - test WRK command generator
func TestWrkBenchCommand(t *testing.T) {
	config := &Config{}
	config.Wrk.Connections = 1
	config.Wrk.Duration = 1
	config.Wrk.Threads = 1
	_, _, err := config.Wrk.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Wrk.Connections = 0
	config.Wrk.Duration = 1
	config.Wrk.Threads = 1
	_, _, err = config.Wrk.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Connections")
	}

	config.Wrk.Connections = 1
	config.Wrk.Duration = 0
	config.Wrk.Threads = 1
	_, _, err = config.Wrk.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Duration")
	}

	config.Wrk.Connections = 1
	config.Wrk.Duration = 1
	config.Wrk.Threads = 0
	_, _, err = config.Wrk.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Threads")
	}
}

// TestSiegeBenchCommand - test Siege command generator
func TestSiegeBenchCommand(t *testing.T) {
	config := &Config{}
	config.Siege.Concurrent = 1
	config.Siege.Time = 1
	_, _, err := config.Siege.BenchCommand("test")
	if err != nil {
		t.Fatal(err)
	}

	config.Siege.Concurrent = 0
	config.Siege.Time = 1
	_, _, err = config.Siege.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Concurrent")
	}

	config.Siege.Concurrent = 1
	config.Siege.Time = 0
	_, _, err = config.Siege.BenchCommand("test")
	if err == nil {
		t.Fatal("Unexpected result for Time")
	}
}
