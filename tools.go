package main

import (
	"fmt"
	"log"
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
	return AB_BENCH + concurency + Keepalive + url
}

// BenchCommand - generate valid WRK command
func (wrk *WrkConfig) BenchCommand(url string) string {
	connections := ""
	if wrk.Connections > 0 {
		connections = fmt.Sprintf(" -c%d ", wrk.Connections)
	} else {
		log.Fatal("wrk connections should be great then 0")
	}
	duration := ""
	if wrk.Duration > 0 {
		duration = fmt.Sprintf(" -d%ds ", wrk.Duration)
	} else {
		log.Fatal("wrk duration should be great then 0")
	}
	threads := ""
	if wrk.Threads > 0 {
		threads = fmt.Sprintf(" -t%d ", wrk.Threads)
	} else {
		log.Fatal("wrk threads should be great then 0")
	}
	return WRK_BENCH + " --latency " + threads + connections + duration + url
}

// BenchCommand - generate valid Siege command
func (s *SiegeConfig) BenchCommand(url string) string {
	concurrent := ""
	if s.Concurrent > 0 {
		concurrent = fmt.Sprintf(" -c%d ", s.Concurrent)
	} else {
		log.Fatal("Siege concurrent should be great then 0")
	}
	time := ""
	if s.Time > 0 {
		time = fmt.Sprintf(" -t%dS ", s.Time)
	} else {
		log.Fatal("Siege time should be great then 0")
	}
	return SIEGE_BENCH + " -b " + concurrent + time + url
}
