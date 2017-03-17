package main

import (
	"flag"
	"fmt"
	cfg "github.com/mrlsd/go-benchmark-app/config"
)

// cliParams - params fo CLI config
var cliParams = &cfg.Config{}

// init - init flags (for Testing it separated)
func init() {
	flag.StringVar(&cfg.ConfigFile, "c", cfg.ConfigFile, "load configuration from `FILE`")
	flag.BoolVar(&cliParams.Verbose, "v", false, "verbose output")
	flag.BoolVar(&cliParams.HtmlOutput, "html", false, "render results to HTML output and open it with default browser")
	flag.Usage = usage
}

// usage - default usaage message
func usage() {
	fmt.Printf("Go Benchmark Applications v%s\nOptions:\n", cfg.AppVersion)
	flag.PrintDefaults()
}

// InitCli - cli flags init and parse
func InitCli() *cfg.Config {
	flag.Parse()
	return cliParams
}
