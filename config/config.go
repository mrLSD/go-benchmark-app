package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

var (
	// ConfigFile - default config file
	ConfigFile = "config/main.toml"
	// AppVersion - current major app version
	AppVersion = "1.0.0"
	// Cfg - global config state
	Cfg = &Config{}
)

const (
	// AB_BENCH - ab benchmark tool
	AB_BENCH = "/usr/bin/ab"
	// WRK_BENCH - wrk benchmark tool
	WRK_BENCH = "/usr/bin/wrk"
	// SIEGE_BENCH - siege benchmark tool
	SIEGE_BENCH = "/usr/bin/siege"
)

// Config - base config
type Config struct {
	Verbose   bool
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
	BinPath    string
	Concurency int
	Keepalive  bool
	Requests   int
}

// WrkConfig - config for WRK benchmark
type WrkConfig struct {
	BinPath     string
	Connections int
	Duration    int
	Threads     int
}

// SiegeConfig - config for Siege benchmark
type SiegeConfig struct {
	BinPath    string
	Concurrent int
	Time       int
}

// AppConfig - configure specific App for bench
// Path should be full valid path to App
type AppConfig struct {
	Title string
	Path  string
	Url   string
}

// LoadConfig - load TOML config file
func LoadConfig(file string, cliParams *Config) (*Config, error) {
	config := cliParams
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return &Config{}, fmt.Errorf("Failed to load config: %s\nReason: %v", file, err)
	}
	Cfg = config
	return config, nil
}
