package main

import (
	"fmt"
)

// BenchCommand - base interface for generating Bench Command
type BenchCommand interface {
	BenchCommand(url string) (string, []string, error)
}

// BenchCommand - generate valid AB command
func (ab *AbConfig) BenchCommand(url string) (string, []string, error) {
	var params []string
	if ab.Keepalive {
		params = append(params, "-k")
	}
	if ab.Concurency > 0 {
		if ab.Concurency > ab.Requests {
			return "", params, fmt.Errorf("ab requests  = %d, should be great or equal concurency = %d", ab.Requests, ab.Concurency)
		}
		params = append(params, "-c", string(ab.Concurency), "-n", string(ab.Requests))
	} else {
		return "", params, fmt.Errorf("ab concurency  = %d, should be great then 0", ab.Concurency)
	}
	params = append(params, url)
	return AB_BENCH, params, nil
}

// BenchCommand - generate valid WRK command
func (wrk *WrkConfig) BenchCommand(url string) (string, []string, error) {
	var params []string
	if wrk.Connections > 0 {
		params = append(params, fmt.Sprintf(" -c%d ", wrk.Connections))
	} else {
		return "", params, fmt.Errorf("wrk connections = %d,  should be great then 0", wrk.Connections)
	}
	if wrk.Duration > 0 {
		params = append(params, fmt.Sprintf(" -d%ds ", wrk.Duration))
	} else {
		return "", params, fmt.Errorf("wrk duration = %d, should be great then 0", wrk.Duration)
	}
	if wrk.Threads > 0 {
		params = append(params, fmt.Sprintf(" -t%d ", wrk.Threads))
	} else {
		return "", params, fmt.Errorf("wrk threads = %d, should be great then 0", wrk.Threads)
	}
	params = append(params, "--latency")
	params = append(params, url)
	return WRK_BENCH, params, nil
}

// BenchCommand - generate valid Siege command
func (s *SiegeConfig) BenchCommand(url string) (string, []string, error) {
	var params []string
	if s.Concurrent > 0 {
		params = append(params, fmt.Sprintf(" -c%d ", s.Concurrent))
	} else {
		return "", params, fmt.Errorf("Siege concurrent = %d, should be great then 0", s.Concurrent)
	}
	if s.Time > 0 {
		params = append(params, fmt.Sprintf(" -t%dS ", s.Time))
	} else {
		return "", params, fmt.Errorf("Siege time = %d, should be great then 0", s.Time)
	}
	params = append(params, "-b")
	params = append(params, url)
	return SIEGE_BENCH, params, nil
}
