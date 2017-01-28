package main

import (
	"fmt"
	"log"
)

// LogFatal - alias for logger, for simplify testing coverage
var LogFatal = log.Fatal

func main() {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		LogFatal(err)
	}
	fmt.Printf("%s\nversion: %s\n", config.Title, config.Version)
	if err := RunBanchmarks(config); err != nil {
		LogFatal(err)
	}
}
