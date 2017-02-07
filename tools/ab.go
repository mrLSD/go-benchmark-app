package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
	"strconv"
	"strings"
)

// AbResults - results for AB benchmarks
type AbResults struct {
	commandResults
	FailedRequests    float64
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
	var err error

	failedRequests := regexp.MustCompile(`Failed.requests:[\s]+([\d]+)`)
	requestsPerSecond := regexp.MustCompile(`Requests.per.second:[\s]+([\d\.]+).\[`)
	timePerRequest := regexp.MustCompile(`Time.per.request:[\s]+([\d\.]+).\[([a-z]+)\].\(mean\)`)
	timePerRequestAll := regexp.MustCompile(`Time.per.request:[\s]+([\d\.]+).\[([a-z]+)\].\(mean\,`)
	transferRate := regexp.MustCompile(`Transfer.rate:[\s]+([\d\.]+).\[(.+)\/.*received`)

	res := failedRequests.FindSubmatch(data)
	if len(res) > 1 {
		result.FailedRequests, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tFailed Requests:\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = requestsPerSecond.FindSubmatch(data)
	if len(res) > 1 {
		result.RequestsPerSecond, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tRequests Per Second:\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = timePerRequest.FindSubmatch(data)
	if len(res) > 2 {
		result.TimePerRequest.Time, _ = strconv.ParseFloat(string(res[1]), 32)
		result.TimePerRequest.Quantor = string(res[2])
		if config.Cfg.Verbose {
			fmt.Printf("\tTime Per Request:\t%v\t%v\n", string(res[1]), string(res[2]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = timePerRequestAll.FindSubmatch(data)
	if len(res) > 2 {
		result.TimePerRequestAll.Time, _ = strconv.ParseFloat(string(res[1]), 32)
		result.TimePerRequestAll.Quantor = string(res[2])
		if config.Cfg.Verbose {
			fmt.Printf("\tTime Per Request [avg]:\t%v %v\n", string(res[1]), string(res[2]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = transferRate.FindSubmatch(data)
	if len(res) > 2 {
		result.TransferRate.Transfer, _ = strconv.ParseFloat(string(res[1]), 32)
		result.TransferRate.Rate = string(res[2])
		if config.Cfg.Verbose {
			fmt.Printf("\tTransfer Rate:\t\t%v %v\n", string(res[1]), string(res[2]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	return result, err
}

// Calculate - mean value of fields
func (ab AbResults) Calculate(data *AbResults) AbResults {
	count := float64(config.Cfg.Try)
	result := AbResults{
		FailedRequests:    ab.FailedRequests/count + data.FailedRequests,
		RequestsPerSecond: ab.RequestsPerSecond/count + data.RequestsPerSecond,
	}

	// Check Rate (Kbyte/sec  byte/sec sec) and
	// concatinate it if not exist
	var resultSplit []string
	resultSplit = strings.Split(data.TransferRate.Rate, "|")
	isNew := true
	for i := 0; i < len(resultSplit); i++ {
		if resultSplit[i] == ab.TransferRate.Rate {
			isNew = false
		}
	}
	result.TransferRate.Rate = ab.TransferRate.Rate
	if isNew {
		if data.TransferRate.Rate != "" {
			result.TransferRate.Rate = data.TransferRate.Rate + "|" + ab.TransferRate.Rate
		} else {
			result.TransferRate.Rate = ab.TransferRate.Rate
		}
	} else {
		result.TransferRate.Rate = data.TransferRate.Rate
	}
	result.TransferRate.Transfer = ab.TransferRate.Transfer/count + data.TransferRate.Transfer

	if !strings.Contains(data.TimePerRequest.Quantor, ab.TimePerRequest.Quantor) {
		result.TimePerRequest.Quantor = data.TimePerRequest.Quantor + "" + ab.TimePerRequest.Quantor
	}
	result.TimePerRequest.Time = ab.TimePerRequest.Time/count + data.TimePerRequest.Time

	if !strings.Contains(data.TimePerRequestAll.Quantor, ab.TimePerRequestAll.Quantor) {
		result.TimePerRequestAll.Quantor = data.TimePerRequestAll.Quantor + "" + ab.TimePerRequestAll.Quantor
	}
	result.TimePerRequestAll.Time = ab.TimePerRequestAll.Time/count + data.TimePerRequestAll.Time

	return result
}

// PrintResults - print fields
func (ab AbResults) PrintResults() {
	fmt.Printf("\tFailed Requests:\t%.2f\n", ab.FailedRequests)
	fmt.Printf("\tRequests Per Second:\t%.2f\n", ab.RequestsPerSecond)
	fmt.Printf("\tTime Per Request:\t%.2f %v\n", ab.TimePerRequest.Time, ab.TimePerRequest.Quantor)
	fmt.Printf("\tTime Per Request [avg]:\t%.2f %v\n", ab.TimePerRequestAll.Time, ab.TimePerRequestAll.Quantor)
	fmt.Printf("\tTransfer Rate:\t\t%.2f %v\n", ab.TransferRate.Transfer, ab.TransferRate.Rate)
}
