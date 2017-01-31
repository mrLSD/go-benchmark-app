package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
)

// WrkResults - results for Wrk benchmarks
type WrkResults struct {
	commandResults
	FailedRequests []int
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

// Analyze - for Wrk parsed results
func (wrk WrkResults) Analyze(data []byte) {
	var LatencyStats = regexp.MustCompile(`Latency[\s]+([\w\.]+)[\s]+([\w\.]+)[\s]+([\w\.]+)`)
	var recSecStats = regexp.MustCompile(`Req\/Sec[\s]+([\w\.]+)[\s]+([\w\.]+)[\s]+([\w\.]+)`)
	var latencyDistribution99pers = regexp.MustCompile(`99%[\s]+([\w\.]+)`)
	var reqSec = regexp.MustCompile(`Requests\/sec:[\s]+([\w\.]+)`)
	_ = LatencyStats
	_ = recSecStats
	_ = latencyDistribution99pers
	_ = reqSec
	/*
		res := LatencyStats.FindSubmatch(data)
		fmt.Printf("\t%v\n\t%v\n\t%v\n", string(res[1]), string(res[2]), string(res[3]))
		res = recSecStats.FindSubmatch(data)
		fmt.Printf("\t%v\n\t%v\n\t%v\n", string(res[1]), string(res[2]), string(res[3]))
		res = latencyDistribution99pers.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))
		res = reqSec.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))
	*/
}
