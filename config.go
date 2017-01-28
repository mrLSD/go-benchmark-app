package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

const (
	// CONFIG_FILE - default config file
	CONFIG_FILE = "config/main.toml"
	// AB_BENCH - ab benchmark tool
	AB_BENCH = "/usr/bin/ab"
	// WRK_BENCH - wrk benchmark tool
	WRK_BENCH = "/usr/bin/wrk"
	// SIEGE_BENCH - siege benchmark tool
	SIEGE_BENCH = "/usr/bin/siege"
)

// Config - base config
type Config struct {
	Title     string
	Version   string
	Delay     time.Duration
	WaitToRun time.Duration
	Try       int
	Ab        AbConfig
	Wrk       WrkConfig
	Siege     SiegeConfig
	App       []AppConfig
}

// AbConfig - config for AB benchmark
type AbConfig struct {
	Concurency int
	Keepalive  bool
	Requests   int
}

// WrkConfig - config for WRK benchmark
type WrkConfig struct {
	Connections int
	Duration    int
	Threads     int
}

// SiegeConfig - config for Siege benchmark
type SiegeConfig struct {
	Concurrent int
	Time       int
}

// AppConfig - configure specific App for bench
// Path should be full valid path to App
type AppConfig struct {
	Title string
	Path  string
	url   []string
}

// LoadConfig - load TOML config file
func LoadConfig(file string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return &Config{}, fmt.Errorf("Failed to load config: %s\nReason: %v", file, err)
	}
	return &config, nil
}
