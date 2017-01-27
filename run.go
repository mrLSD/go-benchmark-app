package main

import (
	"os/exec"
)

// KillProcess - alias for Process.Kill()
// It's used for simplify use and testing code
var KillProcess = killProcrss

// RunCommand - alias fo exec.Command.Output
// execute command and returns its standard output
var RunCommand = runCommand

// RunBanchmars - run all benchmarks
func RunBanchmars(config *Config) error {
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
			return err
		}

		for j := 0; j < len(benchmarkTools); j++ {
			command, err := benchmarkTools[j].tool.BenchCommand("http://localhost:3000/")
			if err != nil {
				return err
			}
			// Run specific bench-tool
			output, err := RunCommand(command)
			if err != nil {
				KillProcess(cmd)
				return err
			}
			println(output)
		}

		if err := KillProcess(cmd); err != nil {
			return err
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
func runCommand(command string) ([]byte, error) {
	return exec.Command(command).Output()
}
