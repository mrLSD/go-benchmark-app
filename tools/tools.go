package tools

// Results - interface for useful Results methods
type Results interface {
	Command() string
	Params() []string
	Parse(data []byte) (Results, error)
}

// commandResults - results for Tool command generation
type commandResults struct {
	command string
	params  []string
}

// BenchCommand - base interface for generating Bench Command
type BenchCommand interface {
	BenchCommand(url string) (Results, error)
}

// BenchResults - complex results of benchamarks tools
type BenchResults struct {
	Ab    AbResults
	Wrk   WrkResults
	Siege SiegeResults
}

// AggreatedResults - aggregated results
type AggreatedResults [][]BenchResults
