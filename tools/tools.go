package tools

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
)

// Results - interface for useful Results methods
type Results interface {
	Command() string
	Params() []string
	Parse(data []byte) (Results, error)
}

// commandResults - results for Tool command generation
type commandResults struct {
	command string
	params  []string
}

// BenchCommand - base interface for generating Bench Command
type BenchCommand interface {
	BenchCommand(url string) (Results, error)
}

// CalculateResults - interface for
// calculation benchmarks results
type CalculateResults interface {
	Calculate(data *BenchResults) BenchResults
}

// BenchResults - complex results of benchamarks tools
type BenchResults struct {
	Ab    AbResults
	Wrk   WrkResults
	Siege SiegeResults
}

// AggreatedResults - aggregated results
type AggreatedResults [][]BenchResults

// Calculate benchmarks results
func (br *BenchResults) Calculate(data *BenchResults) BenchResults {
	return BenchResults{
		Ab:    br.Ab.Calculate(&data.Ab),
		Wrk:   br.Wrk.Calculate(&data.Wrk),
		Siege: br.Siege.Calculate(&data.Siege),
	}
}

// DataAnalyze - analyzy benchdata and print it
func (ar *AggreatedResults) DataAnalyze() []BenchResults {
	results := make([]BenchResults, len(*ar))
	for app := 0; app < len(*ar); app++ {
		for i := 0; i < len((*ar)[app]); i++ {
			results[app] = (*ar)[app][i].Calculate(&results[app])
		}
	}
	return results
}

// PrintResults - to standart output
func PrintResults(results *[]BenchResults, config *config.Config) {
	for i := 0; i < len(*results); i++ {
		fmt.Printf("%s:\n", config.App[i].Title)
		fmt.Println("  [Ab Benchmark results]")
		(*results)[i].Ab.PrintResults()
		fmt.Println("  [Wrk Benchmark results]")
		(*results)[i].Wrk.PrintResults()
		fmt.Println("  [Siege Benchmark results]")
		(*results)[i].Siege.PrintResults()
	}
}
