package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Block interface {
	Execute(p models.Puzzle) Block
}

type BlockGenerator struct{}

func NewBlockGenerator() *BlockGenerator {
	return &BlockGenerator{}
}

func (g *BlockGenerator) Action(moves []string) Block {
	return NewAction(moves)
}

func (g *BlockGenerator) Condition(condFunc ConditionFunc) Block {
	return NewCondition(condFunc)
}
