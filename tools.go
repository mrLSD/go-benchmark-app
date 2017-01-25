package main

import (
	"log"
	"fmt"
)

// BenchCommand - base interface for generating Bench Command
type BenchCommand interface {
	BenchCommand(config *Config)
}

// BenchCommand - generate valid AB command
func (ab *AbConfig) BenchCommand(url string) string {
	Keepalive := ""
	if ab.Keepalive {
		Keepalive = " -k "
	}
	concurency := ""
	if ab.Concurency > 0 {
		if ab.Concurency > ab.Requests {
			log.Fatal("ab requests should be great or equal concurency")
		}
		concurency = fmt.Sprintf(" -c %d -n %d ", ab.Concurency, ab.Requests)

	} else {
		log.Fatal("ab concurency should be great then 0")
	}
	return "ab " + concurency + Keepalive + url
}
