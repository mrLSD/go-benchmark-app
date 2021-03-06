package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"regexp"
	"strconv"
)

// SiegeResults - results for Siege benchmarks
type SiegeResults struct {
	commandResults
	Transactions       float64
	Availability       float64
	TransactionRate    float64
	Concurrency        float64
	LongestTransaction float64
}

// SiegeTool - benchmark tool
type SiegeTool struct {
	*config.SiegeConfig
}

// BenchCommand - generate valid Siege command
func (s SiegeTool) BenchCommand(url string) (Results, error) {
	var params []string
	results := SiegeResults{}
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
	params = append(params, "-b")
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

// Parse - for Siege parsed results
func (s SiegeResults) Parse(data []byte) (Results, error) {
	var result SiegeResults
	var err error

	transactions := regexp.MustCompile(`Transactions:[\s]+([\d\.]+)`)
	availability := regexp.MustCompile(`Availability:[\s]+([\d\.]+)`)
	transactionRate := regexp.MustCompile(`Transaction.rate:[\s]+([\d\.]+)`)
	concurrency := regexp.MustCompile(`Concurrency:[\s]+([\d\.]+)`)
	longestTransaction := regexp.MustCompile(`Longest.transaction:[\s]+([\d\.]+)`)

	res := transactions.FindSubmatch(data)
	if len(res) > 1 {
		result.Transactions, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tTransactions:\t\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = availability.FindSubmatch(data)
	if len(res) > 1 {
		result.Availability, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tAvailability:\t\t%v%%\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = transactionRate.FindSubmatch(data)
	if len(res) > 1 {
		result.TransactionRate, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tTransaction Rate:\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = concurrency.FindSubmatch(data)
	if len(res) > 1 {
		result.Concurrency, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tConcurrency:\t\t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	res = longestTransaction.FindSubmatch(data)
	if len(res) > 1 {
		result.LongestTransaction, _ = strconv.ParseFloat(string(res[1]), 32)
		if config.Cfg.Verbose {
			fmt.Printf("\tLongest Transaction: \t%v\n", string(res[1]))
		}
	} else {
		err = fmt.Errorf("%v\n\tParse error: %v", err, res)
	}

	return result, err
}

// Calculate - mean value of fields
func (s SiegeResults) Calculate(data *SiegeResults) SiegeResults {
	count := float64(config.Cfg.Try)
	result := SiegeResults{
		Transactions:       s.Transactions/count + data.Transactions,
		Availability:       s.Availability/count + data.Availability,
		TransactionRate:    s.TransactionRate/count + data.TransactionRate,
		Concurrency:        s.Concurrency/count + data.Concurrency,
		LongestTransaction: s.LongestTransaction/count + data.LongestTransaction,
	}

	return result
}

// PrintResults - print fields
func (s SiegeResults) PrintResults() {
	fmt.Printf("\tTransactions:\t\t%.2f\n", s.Transactions)
	fmt.Printf("\tAvailability:\t\t%.2f%%\n", s.Availability)
	fmt.Printf("\tTransaction Rate:\t%.2f\n", s.TransactionRate)
	fmt.Printf("\tConcurrency:\t\t%.2f\n", s.Concurrency)
	fmt.Printf("\tLongest Transaction: \t%.2f\n", s.LongestTransaction)
}
