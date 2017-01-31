package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
)

// SiegeResults - results for Siege benchmarks
type SiegeResults struct {
	commandResults
	Transactions       []int
	Availability       []float32
	TransactionRate    []float32
	Concurrency        []float32
	LongestTransaction []float32
}

// SiegeTool - benchmark tool
type SiegeTool struct {
	*config.SiegeConfig
}

// BenchCommand - generate valid Siege command
func (s SiegeTool) BenchCommand(url string) (Results, error) {
	var params []string
	var results SiegeResults = SiegeResults{}
	if s.Concurrent > 0 {
		params = append(params, fmt.Sprintf("-c%d", s.Concurrent))
	} else {
		return results, fmt.Errorf("Siege concurrent = %d, should be great then 0", s.Concurrent)
	}
	if s.Time > 0 {
		params = append(params, fmt.Sprintf("-t%dS", s.Time))
	} else {
		return results, fmt.Errorf("Siege time = %d, should be great then 0", s.Time)
	}
	params = append(params, url)
	results.command = config.SIEGE_BENCH
	results.params = params
	return results, nil
}

// Command - for Siege command tool
func (s SiegeResults) Command() string {
	return s.command
}

// Params - for Siege command tool
func (s SiegeResults) Params() []string {
	return s.params
}

// Analyze - for Siege parsed results
func (s SiegeResults) Analyze(data []byte) {
	var transactions = regexp.MustCompile(`Transactions:[\s]+([\d\.]+)`)
	var availability = regexp.MustCompile(`Availability:[\s]+([\d\.]+)`)
	var transactionRate = regexp.MustCompile(`Transaction.rate:[\s]+([\d\.]+)`)
	var concurrency = regexp.MustCompile(`Concurrency:[\s]+([\d\.]+)`)
	var longestTransaction = regexp.MustCompile(`Longest.transaction:[\s]+([\d\.]+)`)

	_ = transactions
	_ = availability
	_ = transactionRate
	_ = concurrency
	_ = longestTransaction
	/*
		res := transactions.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))

		res = availability.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))

		res = transactionRate.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))

		res = concurrency.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))

		res = longestTransaction.FindSubmatch(data)
		fmt.Printf("\t%v\n", string(res[1]))
	*/
}
