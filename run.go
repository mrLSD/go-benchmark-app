package main

import (
	"os/exec"
)

var KillProcess = killProcrss

// RunBanchmars - run all benchmarks
func RunBanchmars(config *Config) error {
	for i := 0; i < len(config.App); i++ {
		println("===============================")
		println(config.App[i].Title)
		cmd := exec.Command(config.App[i].Path)
		if err := cmd.Start(); err != nil {
			return err
		}

		command, err := config.Ab.BenchCommand("http://localhost:3000/")
		if err != nil {
			return err
		}
		println(command)

		command, err = config.Wrk.BenchCommand("http://localhost:3000/")
		if err != nil {
			return err
		}
		println(command)

		command, err = config.Siege.BenchCommand("http://localhost:3000/")
		if err != nil {
			return err
		}
		println(command)

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
