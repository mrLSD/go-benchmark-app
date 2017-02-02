package tools

import (
	cfg "github.com/mrlsd/go-benchmark-app/config"
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

`

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
	data = []byte(AB_RESULT)
	result.Analyze(data)
}
