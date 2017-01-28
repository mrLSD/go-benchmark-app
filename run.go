package main

import (
	"fmt"
	"os/exec"
	"time"
)

// KillProcess - alias for Process.Kill()
// It's used for simplify use and testing code
var KillProcess = killProcrss

// RunCommand - alias fo exec.Command.Output
// execute command and returns its standard output
var RunCommand = runCommand

// RunBanchmars - run all benchmarks
func RunBanchmarks(config *Config) error {
	// Collect bench-tools to array
	benchmarkTools := []struct {
		tool BenchCommand
	}{
		{tool: &config.Ab},
		{tool: &config.Wrk},
		{tool: &config.Siege},
	}

	for i := 0; i < len(config.App); i++ {
		println("===============================")
		println(config.App[i].Title)
		cmd := exec.Command(config.App[i].Path)
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("Failed execute:\n\t%s\n\t%s", config.App[i].Path, err.Error())
		}
		time.Sleep(config.WaitToRun * time.Second)

		for j := 0; j < len(benchmarkTools); j++ {
			command, params, err := benchmarkTools[j].tool.BenchCommand("http://localhost:3000/")
			if err != nil {
				return fmt.Errorf("Failed run bachmark tool:\n\t%s \n\t%v \n\t%s", command, params, err)
			}
			// Run specific bench-tool
			fmt.Printf("Run command: %s\n", command)
			output, err := RunCommand(command, params...)
			if err != nil {
				KillProcess(cmd)
				println(string(output))
				return fmt.Errorf("Bachmark failed result:\n\t%s \n\t%v \n\t%s", command, params, err)
			}
			time.Sleep(config.Delay * time.Second)
		}

		if err := KillProcess(cmd); err != nil {
			return fmt.Errorf("KillProcess error: %s", err.Error())
		}
	}
	return nil
}

// killProcrss - immediately kill process
func killProcrss(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}

// killProcrss - execute command and
// returns its standard output
func runCommand(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}
