package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
	"strconv"
)

// WrkResults - results for Wrk benchmarks
type WrkResults struct {
	commandResults
	FailedRequests float64
	ReqSec         float64
	Requests       float64
	LatencyStats   struct {
		Avg, Stdev, Max struct {
			Time float64
			// ms, us atc...
			Quantor string
		}
	}
	RecSecStats struct {
		Avg, Stdev, Max struct {
			Transfer float64
			// k, m atc...
			Quantor string
		}
	}
	LatencyDistribution99pers struct {
		Time float64
		// ms, us atc...
		Quantor string
	}
}

// WrkTool - benchmark tool
type WrkTool struct {
	*config.WrkConfig
}

// BenchCommand - generate valid WRK command
func (wrk WrkTool) BenchCommand(url string) (Results, error) {
	var params []string
	var results WrkResults = WrkResults{}
	if wrk.Connections > 0 {
		params = append(params, fmt.Sprintf("-c%d", wrk.Connections))
	} else {
		return results, fmt.Errorf("wrk connections = %d,  should be great then 0", wrk.Connections)
	}
	if wrk.Duration > 0 {
		params = append(params, fmt.Sprintf("-d%ds", wrk.Duration))
	} else {
		return results, fmt.Errorf("wrk duration = %d, should be great then 0", wrk.Duration)
	}
	if wrk.Threads > 0 {
		params = append(params, fmt.Sprintf("-t%d", wrk.Threads))
	} else {
		return results, fmt.Errorf("wrk threads = %d, should be great then 0", wrk.Threads)
	}
	params = append(params, "--latency")
	params = append(params, url)
	results.command = config.WRK_BENCH
	results.params = params
	return results, nil
}

// Command - for Wrk command tool
func (wrk WrkResults) Command() string {
	return wrk.command
}

// Params - for Wrk command tool
func (wrk WrkResults) Params() []string {
	return wrk.params
}

// Parse - for Wrk parsed results
func (wrk WrkResults) Parse(data []byte) (Results, error) {
	var result WrkResults
	var err error

	latencyStats := regexp.MustCompile(`Latency[\s]+([\d\.]+)([\w]+)[\s]+([\d\.]+)([\w]+)[\s]+([\d\.]+)([\w]+)`)
	recSecStats := regexp.MustCompile(`Req\/Sec[\s]+([\d\.]+)([\w]+)[\s]+([\d\.]+)([\w]+)[\s]+([\d\.]+)([\w]+)`)
	latencyDistribution99pers := regexp.MustCompile(`99%[\s]+([\d\.]+)([\w]+)`)
	reqSec := regexp.MustCompile(`Requests\/sec:[\s]+([\w\.]+)`)
	requests := regexp.MustCompile(`[\s]+([\w\.]+)[\s]+requests`)
	failedRequests := regexp.MustCompile(`Non\-2xx[\w\s]+responses:[\s]+([\w\.]+)`)

	res := latencyStats.FindSubmatch(data)
	if len(res) > 6 {
		result.LatencyStats.Avg.Time, _ = strconv.ParseFloat(string(res[1]), 32)
		result.LatencyStats.Avg.Quantor = string(res[2])

		result.LatencyStats.Stdev.Time, _ = strconv.ParseFloat(string(res[3]), 32)
		result.LatencyStats.Stdev.Quantor = string(res[4])

		result.LatencyStats.Max.Time, _ = strconv.ParseFloat(string(res[5]), 32)
		result.LatencyStats.Max.Quantor = string(res[6])

		if config.Cfg.Verbose {
			fmt.Printf("\tLatency Stats Avg:\t%v %v\n", string(res[1]), string(res[2]))
			fmt.Printf("\tLatency Stats Stdev:\t%v %v\n", string(res[3]), string(res[4]))
			fmt.Printf("\tLatency Stats Max:\t%v %v\n", string(res[5]), string(res[6]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = recSecStats.FindSubmatch(data)
	if len(res) > 6 {
		result.RecSecStats.Avg.Transfer, _ = strconv.ParseFloat(string(res[1]), 32)
		result.RecSecStats.Avg.Quantor = string(res[2])

		result.RecSecStats.Stdev.Transfer, _ = strconv.ParseFloat(string(res[3]), 32)
		result.RecSecStats.Stdev.Quantor = string(res[4])

		result.RecSecStats.Max.Transfer, _ = strconv.ParseFloat(string(res[5]), 32)
		result.RecSecStats.Max.Quantor = string(res[6])

		if config.Cfg.Verbose {
			fmt.Printf("\tReq/Sec Stats Avg:\t%v %v\n", string(res[1]), string(res[2]))
			fmt.Printf("\tReq/Sec Stats Stdev:\t%v %v\n", string(res[3]), string(res[4]))
			fmt.Printf("\tReq/Sec Stats Max:\t%v %v\n", string(res[5]), string(res[6]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = latencyDistribution99pers.FindSubmatch(data)
	if len(res) > 2 {
		result.LatencyDistribution99pers.Time, _ = strconv.ParseFloat(string(res[1]), 32)
		result.LatencyDistribution99pers.Quantor = string(res[2])
		if config.Cfg.Verbose {
			fmt.Printf("\tLatency Distribution\n\t\t\t[99%%]:\t%v %v\n", string(res[1]), string(res[2]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = reqSec.FindSubmatch(data)
	if len(res) > 1 {
		result.ReqSec, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tRequests/sec:\t\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = requests.FindSubmatch(data)
	if len(res) > 1 {
		result.Requests, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tRequests:\t\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = failedRequests.FindSubmatch(data)
	if len(res) > 1 {
		result.FailedRequests, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tFailed Requests:\t%v\n", string(res[1]))
		}
	} else {
		if config.Cfg.Verbose {
			fmt.Printf("\tFailed Requests:\t%v\n", 0)
		}
		result.FailedRequests = 0
	}

	return result, err
}
