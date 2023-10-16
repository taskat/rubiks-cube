package panics

import "fmt"

type MaxRunReached struct {
	runs int
}

func NewMaxRunReached(runs int) *MaxRunReached {
	return &MaxRunReached{runs: runs}
}

func (m *MaxRunReached) Error() string {
	return m.String()
}

func (m *MaxRunReached) String() string {
	return fmt.Sprintf("Goal reached max runs (%d)", m.runs)
}
