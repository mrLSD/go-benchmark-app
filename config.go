package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const (
	CONFIG_FILE = "config/main.toml"
	AB_BENCH    = "wrk"
	WRK_BENCH   = "wrk"
	SIEGE_BENCH = "siege"
)

type Config struct {
	Title   string
	Version string
	Delay   int
	Try     int
	Ab      AbConfig
	Wrk     WrkConfig
	Siege   SiegeConfig
	App     []AppConfig
}

type AbConfig struct {
	Concurency int
	Keepalive  bool
	Requests   int
}

type WrkConfig struct {
	Connections int
	Duration    int
	Threads     int
}

type SiegeConfig struct {
	Concurrent int
}

type AppConfig struct {
	Title string
	Path  string
	url   []string
}

func LoadConfig() *Config {
	var config Config
	if _, err := toml.DecodeFile(CONFIG_FILE, &config); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s\nReason: %v", CONFIG_FILE, err))
	}
	return &config
}
