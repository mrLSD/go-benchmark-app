package main

import (
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
	"log"
)

// LogFatal - alias for logger, for simplify testing coverage
var LogFatal = log.Fatal

func main() {
	cliParams := InitCli()
	config, err := cfg.LoadConfig(cfg.ConfigFile, cliParams)
	if err != nil {
		LogFatal(err)
	}
	fmt.Printf("%s\nversion: %s\n", config.Title, config.Version)
	results, err := RunBenchmarks(config)
	if err != nil {
		LogFatal(err)
	}
	DataAnalyze(&results)
}
