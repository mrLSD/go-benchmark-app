package main

import (
	"flag"
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
)

// usage - default usaage message
func usage() {
	fmt.Printf("Go Benchmark Applications v%s\nOptions:\n", cfg.AppVersion)
	flag.PrintDefaults()
}

// InitCli - cli flags init and parse
func InitCli() *cfg.Config {
	cliParams := &cfg.Config{}
	flag.StringVar(&cfg.ConfigFile, "c", cfg.ConfigFile, "load configuration from `FILE`")
	flag.BoolVar(&cliParams.Verbose, "v", false, "verbose output")
	flag.Usage = usage
	flag.Parse()
	return cliParams
}
