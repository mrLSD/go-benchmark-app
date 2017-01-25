package main

import (
	"log"
	"os/exec"
)

func RunBanchmars(config *Config) {
	for i := 0; i < len(config.App); i++ {
		println("===============================")
		println(config.App[i].Title)
		cmd := exec.Command(config.App[i].Path)
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		println(config.Ab.BenchCommand("http://localhost:3000/"))
		cmd.Process.Kill()
	}
}
