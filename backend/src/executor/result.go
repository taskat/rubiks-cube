package executor

type ExecutionResult struct {
	Turns map[string][]string
	Steps []string
}

func NewExecutionResult(turns map[string][]string, steps []string) *ExecutionResult {
	return &ExecutionResult{Turns: turns, Steps: steps}
}
