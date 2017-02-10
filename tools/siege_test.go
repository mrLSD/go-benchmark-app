package tools

import (
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

const SIEGE_RESULT = `
Transactions:		      146229 hits
Availability:		       99.63 %
Elapsed time:		        9.11 secs
Data transferred:	        2.65 MB
Response time:		        0.01 secs
Transaction rate:	    16051.48 trans/sec
Throughput:		        0.29 MB/sec
Concurrency:		       91.92
Successful transactions:           0
Failed transactions:	         546
Longest transaction:	        0.31
Shortest transaction:	        0.00

`

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

	cfg.Cfg.Verbose = true
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
	data := []byte("")
	result.Parse(data)
	data = []byte(SIEGE_RESULT)
	result.Parse(data)
}

// TestSiegeCalculate - test Siege total results calculation
func TestSiegeCalculate(t *testing.T) {
	initConfig := &cfg.Config{}
	_, err := cfg.LoadConfig("../"+cfg.ConfigFile, initConfig)
	if err != nil {
		t.Fatal(err)
	}
	cfg.Cfg.Try = 3

	// Init Aggregated results
	data := make(AggreatedResults, 1)
	data[0] = make([]BenchResults, cfg.Cfg.Try)

	// Init Results 1
	result1 := SiegeResults{}
	result1.Transactions = 90.
	result1.Availability = 99.
	result1.Concurrency = 90.
	result1.LongestTransaction = 1.2
	result1.TransactionRate = 1200

	// Init Results 2
	result2 := SiegeResults{}
	result2.Transactions = 180.
	result2.Availability = 66.
	result2.Concurrency = 60.
	result2.LongestTransaction = 1.5
	result2.TransactionRate = 1500

	data[0][0].Siege = result1
	data[0][1].Siege = result2
	data[0][2].Siege = result2

	result := data.DataAnalyze()
	if len(result) > 1 {
		t.Fatalf("Faile result length: %v", "DataAnalyze")
	}

	// Test PrintResults
	result[0].Siege.PrintResults()

	if int(result[0].Siege.Transactions) != 150 {
		t.Fatalf("Error calculation: %v", "Siege.TransactionRate")
	}

	if int(result[0].Siege.Availability) != 77 {
		t.Fatalf("Error calculation: %v", "Siege.Availability")
	}

	if result[0].Siege.LongestTransaction != 1.4 {
		t.Fatalf("Error calculation: %v", "Siege.LongestTransaction")
	}

	if int(result[0].Siege.TransactionRate) != 1400 {
		t.Fatalf("Error calculation: %v", "Siege.TransactionRate")
	}

	config := &cfg.Config{
		App: []cfg.AppConfig{
			{
				Title: "Test 1",
				Path:  "/bin/bash",
				Url:   "test",
			},
		},
	}
	PrintResults(&result, config)
}
