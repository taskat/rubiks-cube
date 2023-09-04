package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Block interface {
	Execute() Block
}

type BlockGenerator struct {
	puzzle models.Puzzle
}

func NewBlockGenerator(puzzle models.Puzzle) *BlockGenerator {
	return &BlockGenerator{puzzle: puzzle}
}

func (g *BlockGenerator) Action(next Block, moves []string) Block {
	return newAction(g.puzzle, next, moves)
}

func (g *BlockGenerator) Condition(trueBlock, falseBlock Block, condFunc ConditionFunc) Block {
	return newCondition(g.puzzle, trueBlock, falseBlock, condFunc)
}
