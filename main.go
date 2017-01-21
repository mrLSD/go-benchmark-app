package main

import "fmt"

func main() {
	config := LoadConfig()
	fmt.Printf("%s\nversion: %s\n", config.Title, config.Version)
}