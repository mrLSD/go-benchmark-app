package main

import (
	"fmt"
	"log"
)

var LogFatal = log.Fatal

func main() {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		LogFatal(err)
	}
	fmt.Printf("%s\nversion: %s\n", config.Title, config.Version)
	if err := RunBanchmars(config); err != nil {
		LogFatal(err)
	}
}
