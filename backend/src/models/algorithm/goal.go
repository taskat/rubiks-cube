package algorithm

import (
	"fmt"

	"github.com/taskat/rubiks-cube/src/models"
)

type Goal struct {
	*Condition
	runs    int
	maxRuns int
}

func NewGoal(condFunc ConditionFunc, runs int) *Goal {
	return &Goal{Condition: NewCondition(condFunc), runs: 0, maxRuns: runs}
}

func (g *Goal) Execute(p models.Puzzle) Block {
	if g.runs > g.maxRuns {
		panic("Goal reached max runs" + fmt.Sprintf(" (%d)", g.maxRuns))
	}
	g.runs++
	return g.Condition.Execute(p)
}
