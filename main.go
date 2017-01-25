package main

import "fmt"

func main() {
	config := LoadConfig(CONFIG_FILE)
	fmt.Printf("%s\nversion: %s\n", config.Title, config.Version)

	RunBanchmars(config)
}
