package main

import (
	"os/exec"
)

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
		if err := cmd.Process.Kill(); err != nil {
			return err
		}
	}
	return nil
}
