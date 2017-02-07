package main

import (
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const AB_RESULT = `
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Finished 300 requests


Server Software:
Server Hostname:        localhost
Server Port:            3000

Document Path:          /123
Document Length:        19 bytes

Concurrency Level:      100
Time taken for tests:   0.027 seconds
Complete requests:      300
Failed requests:        0
Non-2xx responses:      300
Total transferred:      38400 bytes
HTML transferred:       5700 bytes
Requests per second:    11038.75 [#/sec] (mean)
Time per request:       9.059 [ms] (mean)
Time per request:       0.091 [ms] (mean, across all concurrent requests)
Transfer rate:          1379.84 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    3   0.9      3       5
Processing:     1    5   1.7      5       8
Waiting:        1    3   1.5      4       8
Total:          5    8   1.7      7      13

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%      8
  80%      9
  90%     10
  95%     11
  98%     12
  99%     13
 100%     13 (longest request)
==>> SIEGE
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

// Alias for success runned command
var runCommandSuccess = func(c string, args ...string) ([]byte, error) {
	println("\n	=> runCommandSuccess")
	switch c {
	case cfg.AB_BENCH:
		return []byte(AB_RESULT), nil
	case cfg.WRK_BENCH:
		return []byte(WRK_RESULT), nil
	case cfg.SIEGE_BENCH:
		return []byte(SIEGE_RESULT), nil
	}
	return []byte(""), nil
}

// Alias for success comand  execution
// but wrong output results
var runCommandSuccessButWronOutput = func(c string, args ...string) ([]byte, error) {
	println("	=> runCommandSuccessButWronOutput")
	return []byte("test"), nil
}

// Alias for failed runned command
var runCommandFailed = func(c string, args ...string) ([]byte, error) {
	println("	=> runCommandFailed")
	return []byte("test"), fmt.Errorf("test %s", "test")
}

// TestRunBenchmarks - with basic cinfig
func TestRunBenchmarks(t *testing.T) {
	RunBenchmarks = runBenchmarks
	initConfig := &cfg.Config{}
	config, err := cfg.LoadConfig(cfg.ConfigFile, initConfig)
	if err != nil {
		t.Fatal(err)
	}

	// Simple check for runCommand
	_, _ = runCommand("/bin/bash")

	RunCommand = runCommandSuccess
	config.Try = 1
	config.WaitToRun = 0
	config.Delay = 0

	_, err = RunBenchmarks(config)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			t.Fatal(err)
		}
	}
}

// TestRunBenchmarksWithWrongAppPath - test with basic config
// and wrong App Path
func TestRunBenchmarksWithWrongAppPath(t *testing.T) {
	RunBenchmarks = runBenchmarks
	initConfig := &cfg.Config{}
	config, err := cfg.LoadConfig(cfg.ConfigFile, initConfig)
	if err != nil {
		t.Fatal(err)
	}

	config.Try = 1
	config.WaitToRun = 0
	config.Delay = 0

	if len(config.App) > 0 {
		config.App[0].Path = "test/test"
	} else {
		t.Fatal("You should have at least one App")
	}
	_, err = RunBenchmarks(config)
	if err == nil {
		t.Fatal("Unexpected exec start result")
	}
}

// TestRunBenchmarksWithWrongParams - with basic config
// but some wrong params
func TestRunBenchmarksWithWrongParams(t *testing.T) {
	RunBenchmarks = runBenchmarks
	initConfig := &cfg.Config{}
	config, err := cfg.LoadConfig(cfg.ConfigFile, initConfig)
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

	config.Try = 1
	config.WaitToRun = 0
	config.Delay = 0

	// Re-init app
	config.App = []cfg.AppConfig{{
		Title: "Test Bash",
		Path:  "/bin/bash",
	}}

	// Success benchmarks
	RunCommand = runCommandSuccess
	_, err = RunBenchmarks(config)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal(err)
		}
	}

	// Failed benchmarks
	RunCommand = runCommandFailed
	_, err = RunBenchmarks(config)
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
	_, err = RunBenchmarks(&abConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for abConfig")
		}
	}

	// Wrong WRK Connections parameter
	wrkConfig := *config
	wrkConfig.Wrk.Connections = 0
	_, err = RunBenchmarks(&wrkConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for wrkConfig")
		}
	}

	// Run Verbose
	cfg.Cfg.Verbose = true
	// Wrong Siege Concurrent parameter
	siegeConfig := *config
	siegeConfig.Siege.Concurrent = 0
	_, err = RunBenchmarks(&siegeConfig)
	if err == nil {
		_, ok := err.(*os.PathError)
		if !ok {
			t.Fatal("Unexpected exec for siegeConfig")
		}
	}
	//cfg.Cfg.Verbose = false

	// Simulate Wrong Kill Process
	KillProcess = func(cmd *exec.Cmd) error {
		return fmt.Errorf("test %s", "test")
	}

	_, err = RunBenchmarks(config)
	if err == nil {
		t.Fatal("Unexpected exec for KillProcess")
	}
}

// TestRunBenchmarksWrongParse - test wron output results
// for results parsing test
func TestRunBenchmarksWrongParse(t *testing.T) {
	RunBenchmarks = runBenchmarks
	initConfig := &cfg.Config{}
	config, err := cfg.LoadConfig(cfg.ConfigFile, initConfig)
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

	config.Try = 1
	config.WaitToRun = 0
	config.Delay = 0

	// Re-init app
	config.App = []cfg.AppConfig{{
		Title: "Test Bash",
		Path:  "/bin/bash",
	}}

	RunCommand = runCommandSuccessButWronOutput
	config.Try = 1
	config.WaitToRun = 0
	config.Delay = 0

	_, err = RunBenchmarks(config)
	if err == nil {
		t.Fatal(fmt.Errorf("Unexpected result for: %s", "runCommandSuccessButWronOutput"))
	}
}
