package main

import (
	"github.com/mrlsd/go-benchmark-app/tools"
)

// DataAnalyze - analyzy benchdata and print it
func DataAnalyze(data *tools.AggreatedResults) {
	results := make([]tools.BenchResults, len(*data))
	for app := 0; app < len(*data); app++ {
		for i := 0; i < len((*data)[app]); i++ {
			results[app] = (*data)[app][i].Calculate(&results[app])
		}
	}

	for i := 0; i < len(results); i++ {
		results[i].Ab.PrintResults()
	}
}
