package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\nversion: %s\n", config.Title, config.Version)
	if err := RunBanchmars(config); err != nil {
		log.Fatal(err)
	}
}
