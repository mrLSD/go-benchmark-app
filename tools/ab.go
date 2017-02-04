package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
	"strconv"
)

// AbResults - results for AB benchmarks
type AbResults struct {
	commandResults
	FailedRequests    int
	RequestsPerSecond float64
	TransferRate      struct {
		Transfer float64
		// Kbyte/sec atc...
		Rate string
	}
	TimePerRequest, TimePerRequestAll struct {
		// Time per Request
		Time float64
		// Quantor - sec, min, hour
		Quantor string
	}
}

// AbTool - benchmark tool
type AbTool struct {
	*config.AbConfig
}

// BenchCommand - generate valid AB command
func (ab AbTool) BenchCommand(url string) (Results, error) {
	var params []string
	var results AbResults = AbResults{}

	if ab.Keepalive {
		params = append(params, "-k")
	}
	if ab.Concurency > 0 {
		if ab.Concurency > ab.Requests {
			return results, fmt.Errorf("ab requests  = %d, should be great or equal concurency = %d", ab.Requests, ab.Concurency)
		}
		params = append(params, "-c", fmt.Sprintf("%d", ab.Concurency), "-n", fmt.Sprintf("%d", ab.Requests))
	} else {
		return results, fmt.Errorf("ab concurency  = %d, should be great then 0", ab.Concurency)
	}
	results.command = config.AB_BENCH
	results.params = append(params, url)
	return results, nil
}

// Command - for AB command tool
func (ab AbResults) Command() string {
	return ab.command
}

// Params - for AB command tool
func (ab AbResults) Params() []string {
	return ab.params
}

// Parse - for AB parsed results
func (ab AbResults) Parse(data []byte) (Results, error) {
	var result AbResults
	var err error = nil

	var failedRequests = regexp.MustCompile(`Failed.requests:[\s]+([\d]+)`)
	var requestsPerSecond = regexp.MustCompile(`Requests.per.second:[\s]+([\d\.]+).\[`)
	var timePerRequest = regexp.MustCompile(`Time.per.request:[\s]+([\d\.]+).\[([a-z]+)\].\(mean\)`)
	var timePerRequestAll = regexp.MustCompile(`Time.per.request:[\s]+([\d\.]+).\[([a-z]+)\].\(mean\,`)
	var transferRate = regexp.MustCompile(`Transfer.rate:[\s]+([\d\.]+).\[(.+)\/.*received`)
	_ = failedRequests
	_ = requestsPerSecond
	_ = timePerRequest
	_ = timePerRequestAll
	_ = transferRate
	res := failedRequests.FindSubmatch(data)
	if len(res) > 1 {
		result.FailedRequests, err = strconv.Atoi(string(res[1]))
		if err != nil {
			return result, err
		}
		fmt.Printf("\t%v\n", string(res[1]))
	} else {
		return result, fmt.Errorf("Parse error %v", res)
	}

	res = requestsPerSecond.FindSubmatch(data)
	if len(res) > 1 {
		result.RequestsPerSecond, err = strconv.ParseFloat(string(res[1]), 32)
		if err != nil {
			return result, err
		}
		fmt.Printf("\t%v\n", string(res[1]))
	} else {
		return result, fmt.Errorf("Parse error %v", res)
	}

	res = timePerRequest.FindSubmatch(data)
	if len(res) > 2 {
		fmt.Printf("\t%v\t%v\n", string(res[1]), string(res[2]))
	} else {
		return result, fmt.Errorf("Parse error %v", res)
	}

	res = timePerRequestAll.FindSubmatch(data)
	if len(res) > 2 {
		fmt.Printf("\t%v\t%v\n", string(res[1]), string(res[2]))
	} else {
		return result, fmt.Errorf("Parse error %v", res)
	}

	res = transferRate.FindSubmatch(data)
	if len(res) > 2 {
		fmt.Printf("\t%v\t%v\n", string(res[1]), string(res[2]))
	} else {
		return result, fmt.Errorf("Parse error %v", res)
	}

	return result, err
}
