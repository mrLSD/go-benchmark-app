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
