package algorithm

import (
	"fmt"

	"github.com/taskat/rubiks-cube/src/models"
)

type Goal struct {
	*Condition
	runs     int
	maxRuns  int
	prepares []*Prepare
}

func NewGoal(condFunc ConditionFunc, runs int) *Goal {
	return &Goal{Condition: NewCondition(condFunc), runs: 0, maxRuns: runs, prepares: make([]*Prepare, 0)}
}

func (g *Goal) AddPrepare(p *Prepare) {
	g.prepares = append(g.prepares, p)
}

func (g *Goal) Execute(p models.Puzzle) Block {
	if g.runs > g.maxRuns {
		panic("Goal reached max runs" + fmt.Sprintf(" (%d)", g.maxRuns))
	}
	g.runs++
	for _, prepare := range g.prepares {
		prepare.Start()
	}
	return g.Condition.Execute(p)
}
