package tools

// Results - interface for useful Results methods
type Results interface {
	Command() string
	Params() []string
	Analize()
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
