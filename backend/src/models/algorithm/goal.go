package algorithm

import (
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/panics"
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
		panic(panics.NewMaxRunReached(g.maxRuns))
	}
	g.runs++
	for _, prepare := range g.prepares {
		prepare.Start()
	}
	return g.Condition.Execute(p)
}
