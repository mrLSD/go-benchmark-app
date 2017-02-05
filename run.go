package main

import (
	"fmt"
	"github.com/mrlsd/go-benchmark-app/config"
	"github.com/mrlsd/go-benchmark-app/tools"
	"os/exec"
	"time"
)

// KillProcess - alias for Process.Kill()
// It's used for simplify use and testing code
var KillProcess = killProcess

// RunCommand - alias fo exec.Command.Output
// execute command and returns its standard output
var RunCommand = runCommand

// RunBenchmarks - run all benchmarks
func RunBenchmarks(config *config.Config) error {
	// Init results array
	var benchResults tools.AggreatedResults
	benchResults = make(tools.AggreatedResults, len(config.App))
	for i := 0; i < len(config.App); i++ {
		benchResults[i] = make([]tools.BenchResults, config.Try)
	}

	// Collect bench-tools to array
	benchmarkTools := []struct {
		tool tools.BenchCommand
	}{
		{tool: tools.AbTool{&config.Ab}},
		{tool: tools.WrkTool{&config.Wrk}},
		{tool: tools.SiegeTool{&config.Siege}},
	}

	// Repeate benchmarks
	for repeat := 0; repeat < config.Try; repeat++ {
		//  Go through applications
		for i := 0; i < len(config.App); i++ {
			println("===============================")
			fmt.Printf("%s [%d]", config.App[i].Title, repeat)
			// Get app command and Run it
			cmd := exec.Command(config.App[i].Path)
			if err := cmd.Start(); err != nil {
				return fmt.Errorf("Failed execute:\n\t%s\n\t%s", config.App[i].Path, err.Error())
			}
			// Wait when app starting
			time.Sleep(config.WaitToRun * time.Second)

			// Go through Benchmark tools
			for j := 0; j < len(benchmarkTools); j++ {
				// Generate bench-command
				results, err := benchmarkTools[j].tool.BenchCommand("http://localhost:3000/test")
				if err != nil {
					return fmt.Errorf("Failed run bachmark tool:\n\t%s \n\t%v \n\t%s", results.Command(), results.Params(), err)
				}
				// Run specific bench-tool
				fmt.Printf("Run command: %s\n", results.Command())
				output, err := RunCommand(results.Command(), results.Params()...)
				if err != nil {
					KillProcess(cmd)
					println(string(output))
					return fmt.Errorf("Bachmark failed result:\n\t%s \n\t%v \n\t%s", results.Command(), results.Params(), err)
				}

				fmt.Printf("\t\t%#v\n", err)
				fmt.Printf("\t\t%#v\n", results.Params())

				// Parse bench-output
				parsed, err := results.Parse(output)
				if err != nil {
					return err
				}

				// Aggregate benchmark results by:
				// Application iterator, Repeat iterator, Bench-tool type
				aggregateResults(&parsed, &benchResults[i][repeat])
				time.Sleep(config.Delay * time.Second)
			}

			if err := KillProcess(cmd); err != nil {
				return fmt.Errorf("KillProcess error: %s", err.Error())
			}
		}
	}
	return nil
}

// killProcess - immediately kill process
func killProcess(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}

// runCommand - execute command and
// returns its  output
func runCommand(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}

// aggregateResults - get Bench Resuls by it type
func aggregateResults(data *tools.Results, benchResults *tools.BenchResults) {
	_ = benchResults
    switch values := (*data).(type) {
	case tools.AbResults:
		benchResults.Ab = values
		fmt.Println("\t=> tools.AbResults")
    case tools.WrkResults:
		benchResults.Wrk = values
        fmt.Println("\t=> tools.WrkResults")
	case tools.SiegeResults:
		benchResults.Siege = values
        fmt.Println("\t=> tools.SiegeResults")
	}
}
