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
		for i := 0; i < len(config.App); i++ {
			println("===============================")
			println(config.App[i].Title)
			cmd := exec.Command(config.App[i].Path)
			if err := cmd.Start(); err != nil {
				return fmt.Errorf("Failed execute:\n\t%s\n\t%s", config.App[i].Path, err.Error())
			}
			time.Sleep(config.WaitToRun * time.Second)

			for j := 0; j < len(benchmarkTools); j++ {
				results, err := benchmarkTools[j].tool.BenchCommand("http://localhost:3000/")
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

				results.Analyze(output)
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
