package tools

import (
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"testing"
)

const WRK_RESULT = `
Running 10s test @ http://localhost:3000/123
  100 threads and 5000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   443.73us  701.77us  24.24ms   93.93%
    Req/Sec    29.50k     3.64k   41.13k    73.04%
  Latency Distribution
     50%  269.00us
     75%  483.00us
     90%    0.86ms
     99%    4.21ms
  599207 requests in 10.10s, 73.15MB read
  Non-2xx or 3xx responses: 599207
Requests/sec:  59353.58
Transfer/sec:      7.25MB
`

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

	cfg.Cfg.Verbose = true
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
	data := []byte("")
	result.Parse(data)
	data = []byte(WRK_RESULT)
	result.Parse(data)
}

// TestWrkCalculate - test Wrk total results calculation
func TestWrkCalculate(t *testing.T) {
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
	result1 := WrkResults{}
	result1.FailedRequests = 100.
	result1.ReqSec = 200.
	result1.Requests = 200.
	result1.LatencyStats.Avg.Time = 60.
	result1.LatencyStats.Avg.Quantor = "us"
	result1.LatencyStats.Stdev.Time = 0.12
	result1.LatencyStats.Stdev.Quantor = "ms"
	result1.LatencyStats.Max.Time = 120
	result1.LatencyStats.Max.Quantor = "ms"
	result1.RecSecStats.Avg.Transfer = 60.
	result1.RecSecStats.Avg.Quantor = "k"
	result1.RecSecStats.Stdev.Transfer = 120.
	result1.RecSecStats.Stdev.Quantor = "b"
	result1.RecSecStats.Max.Transfer = 0.6
	result1.RecSecStats.Max.Quantor = "m"
	result1.LatencyDistribution99pers.Time = 1.2
	result1.LatencyDistribution99pers.Quantor = "ms"

	// Init Results 2
	result2 := WrkResults{}
	result2.FailedRequests = 250.
	result2.ReqSec = 350.
	result2.Requests = 350.
	result2.LatencyStats.Avg.Time = 1.2
	result2.LatencyStats.Avg.Quantor = "ms"
	result2.LatencyStats.Stdev.Time = 2400
	result2.LatencyStats.Stdev.Quantor = "us"
	result2.LatencyStats.Max.Time = 1.2
	result2.LatencyStats.Max.Quantor = "s"
	result2.RecSecStats.Avg.Transfer = 12
	result2.RecSecStats.Avg.Quantor = "k"
	result2.RecSecStats.Stdev.Transfer = 2.4
	result2.RecSecStats.Stdev.Quantor = "k"
	result2.RecSecStats.Max.Transfer = 120
	result2.RecSecStats.Max.Quantor = "k"
	result2.LatencyDistribution99pers.Time = 600
	result2.LatencyDistribution99pers.Quantor = "us"

	data[0][0].Wrk = result1
	data[0][1].Wrk = result2
	data[0][2].Wrk = result2

	result := data.DataAnalyze()
	if len(result) > 1 {
		t.Fatalf("Faile result length: %v", "DataAnalyze")
	}

	// Test PrintResults
	result[0].Wrk.PrintResults()
	if int(result[0].Wrk.FailedRequests) != 200 {
		t.Fatalf("Error calculation: %v", "FailedRequests")
	}

	if int(result[0].Wrk.ReqSec) != 300 {
		t.Fatalf("Error calculation: %v", "ReqSec")
	}

	if int(result[0].Wrk.Requests) != 300 {
		t.Fatalf("Error calculation: %v", "Requests")
	}

	if int(result[0].Wrk.LatencyStats.Avg.Time) != 820 {
		t.Fatalf("Error calculation: %v", result[0].Wrk.LatencyStats.Avg.Time)
	}

	if result[0].Wrk.LatencyStats.Avg.Quantor != "us" {
		t.Fatalf("Error calculation: %v", result[0].Wrk.LatencyStats.Avg.Quantor)
	}

	if int(result[0].Wrk.LatencyStats.Stdev.Time) != 1640 {
		t.Fatalf("Error calculation: %v", "LatencyStats.Stdev.Time")
	}

	if result[0].Wrk.LatencyStats.Stdev.Quantor != "us" {
		t.Fatalf("Error calculation: %v", "Wrk.LatencyStats.Stdev.Quantor")
	}

	if int(result[0].Wrk.LatencyStats.Max.Time) != 840 {
		t.Fatalf("Error calculation: %v", "Wrk.LatencyStats.Max.Time")
	}

	if result[0].Wrk.LatencyStats.Max.Quantor != "ms" {
		t.Fatalf("Error calculation: %v", "LatencyStats.Max.Quantor")
	}

	if int(result[0].Wrk.RecSecStats.Avg.Transfer) != 28 {
		t.Fatalf("Error calculation: %v", "Wrk.RecSecStats.Avg.Transfer")
	}

	if result[0].Wrk.RecSecStats.Avg.Quantor != "k" {
		t.Fatalf("Error calculation: %v", "Wrk.RecSecStats.Avg.Quantor")
	}

	if int(result[0].Wrk.RecSecStats.Stdev.Transfer) != 1640 {
		t.Fatalf("Error calculation: %v", "Wrk.RecSecStats.Stdev.Transfer")
	}

	if result[0].Wrk.RecSecStats.Stdev.Quantor != "b" {
		t.Fatalf("Error calculation: %v", "Wrk.RecSecStats.Stdev.Quantor")
	}

	if int(result[0].Wrk.RecSecStats.Max.Transfer) != 280 {
		t.Fatalf("Error calculation: %v", "Wrk.RecSecStats.Max.Transfer")
	}

	if result[0].Wrk.RecSecStats.Max.Quantor != "k" {
		t.Fatalf("Error calculation: %v", "Wrk.RecSecStats.Max.Quantor")
	}

	if int(result[0].Wrk.LatencyDistribution99pers.Time) != 800 {
		t.Fatalf("Error calculation: %v", "Wrk.LatencyDistribution99pers.Time")
	}

	if result[0].Wrk.LatencyDistribution99pers.Quantor != "us" {
		t.Fatalf("Error calculation: %v", "Wrk.LatencyDistribution99pers.Quantor")
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
