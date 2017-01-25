package main

import (
	"fmt"
)

// BenchCommand - base interface for generating Bench Command
type BenchCommand interface {
	BenchCommand(url string) (string, error)
}

// BenchCommand - generate valid AB command
func (ab *AbConfig) BenchCommand(url string) (string, error) {
	Keepalive := ""
	if ab.Keepalive {
		Keepalive = " -k "
	}
	concurency := ""
	if ab.Concurency > 0 {
		if ab.Concurency > ab.Requests {
			return "", fmt.Errorf("ab requests  = %d, should be great or equal concurency = %d", ab.Requests, ab.Concurency)
		}
		concurency = fmt.Sprintf(" -c %d -n %d ", ab.Concurency, ab.Requests)
	} else {
		return "", fmt.Errorf("ab concurency  = %d, should be great then 0", ab.Concurency)
	}
	return AB_BENCH + concurency + Keepalive + url, nil
}

// BenchCommand - generate valid WRK command
func (wrk *WrkConfig) BenchCommand(url string) (string, error) {
	connections := ""
	if wrk.Connections > 0 {
		connections = fmt.Sprintf(" -c%d ", wrk.Connections)
	} else {
		return "", fmt.Errorf("wrk connections = %d,  should be great then 0", wrk.Connections)
	}
	duration := ""
	if wrk.Duration > 0 {
		duration = fmt.Sprintf(" -d%ds ", wrk.Duration)
	} else {
		return "", fmt.Errorf("wrk duration = %d, should be great then 0", wrk.Duration)
	}
	threads := ""
	if wrk.Threads > 0 {
		threads = fmt.Sprintf(" -t%d ", wrk.Threads)
	} else {
		return "", fmt.Errorf("wrk threads = %d, should be great then 0", wrk.Threads)
	}
	return WRK_BENCH + " --latency " + threads + connections + duration + url, nil
}

// BenchCommand - generate valid Siege command
func (s *SiegeConfig) BenchCommand(url string) (string, error) {
	concurrent := ""
	if s.Concurrent > 0 {
		concurrent = fmt.Sprintf(" -c%d ", s.Concurrent)
	} else {
		return "", fmt.Errorf("Siege concurrent = %d, should be great then 0", s.Concurrent)
	}
	time := ""
	if s.Time > 0 {
		time = fmt.Sprintf(" -t%dS ", s.Time)
	} else {
		return "", fmt.Errorf("Siege time = %d, should be great then 0", s.Time)
	}
	return SIEGE_BENCH + " -b " + concurrent + time + url, nil
}
