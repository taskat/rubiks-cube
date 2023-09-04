package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type ConditionFunc func(models.Puzzle) bool

type condition struct {
	puzzle     models.Puzzle
	trueBlock  Block
	falseBlock Block
	condFunc   ConditionFunc
}

func newCondition(puzzle models.Puzzle, trueBlock, falseBlock Block, condFunc ConditionFunc) Block {
	return &condition{puzzle: puzzle, trueBlock: trueBlock, falseBlock: falseBlock, condFunc: condFunc}
}

func (c *condition) Execute() Block {
	if c.condFunc(c.puzzle) {
		return c.trueBlock
	}
	return c.falseBlock
}
