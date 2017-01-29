package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
)

// AbResults - results for AB benchmarks
type AbResults struct {
	commandResults
	FailedRequests    []int
	RequestsPerSecond []float32
	TransferRate      []struct {
		Transfer float32
		// Kbyte/sec atc...
		Rate string
	}
	TtimePerRequest, TimePerRequestAll []struct {
		// Time per Request
		Time float32
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

// Analize - for AB parsed results
func (ab AbResults) Analize() {
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
	/*
		res := failedRequests.FindSubmatch(output)[1]
		fmt.Printf("\t%v\n", string(res))

		res1 := requestsPerSecond.FindSubmatch(output)[1]
		fmt.Printf("\t%v\n", string(res1))

		res1 = timePerRequest.FindSubmatch(output)[1]
		res2 := timePerRequest.FindSubmatch(output)[2]
		fmt.Printf("\t%v\t%v\n", string(res1), string(res2))

		res1 = timePerRequestAll.FindSubmatch(output)[1]
		res2 = timePerRequestAll.FindSubmatch(output)[2]
		fmt.Printf("\t%v\t%v\n", string(res1), string(res2))

		res1 = transferRate.FindSubmatch(output)[1]
		res2 = transferRate.FindSubmatch(output)[2]
		fmt.Printf("\t%v\t%v\n", string(res1), string(res2))
	*/
}
