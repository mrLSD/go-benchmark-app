package main

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

type Config struct {
	Title   string
	Version string
}

func LoadConfig() *Config {
	var config Config
	if _, err := toml.DecodeFile("config/main.toml", &config); err != nil {
		panic(fmt.Sprintf("Faile to lad config: %v", err))
	}
	return &config
}
