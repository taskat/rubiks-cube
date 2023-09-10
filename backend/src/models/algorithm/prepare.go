package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Prepare struct {
	*Action
	maxConsecs int
	consecs    int
	reached    bool
	next       Block
}

func NewPrepare(action *Action, maxConsecs int) *Prepare {
	return &Prepare{Action: action, maxConsecs: maxConsecs}
}

func (p *Prepare) Execute(puzzle models.Puzzle) Block {
	if p.consecs >= p.maxConsecs {
		return p.next
	}
	p.consecs++
	p.reached = true
	return p.Action.Execute(puzzle)
}

func (p *Prepare) NextSetter() Setter {
	return func(block Block) {
		p.next = block
	}
}

func (p *Prepare) Start() {
	if !p.reached {
		p.consecs = 0
	}
	p.reached = false
}
