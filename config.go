package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
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
}

func LoadConfig() *Config {
	var config Config
	if _, err := toml.DecodeFile("config/main.toml", &config); err != nil {
		panic(fmt.Sprintf("Faile to lad config: %v", err))
	}
	return &config
}
