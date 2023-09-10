package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type ConditionFunc func(models.Puzzle) bool

type Condition struct {
	trueBlock  Block
	falseBlock Block
	condFunc   ConditionFunc
}

func NewCondition(condFunc ConditionFunc) *Condition {
	return &Condition{trueBlock: nil, falseBlock: nil, condFunc: condFunc}
}

func (c *Condition) Execute(p models.Puzzle) Block {
	if c.condFunc(p) {
		return c.trueBlock
	}
	return c.falseBlock
}

func (c *Condition) FalseSetter() Setter {
	return func(block Block) {
		c.falseBlock = block
	}
}

func (c *Condition) TrueSetter() Setter {
	return func(block Block) {
		c.trueBlock = block
	}
}
