package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Action struct {
	next  Block
	moves []string
}

func NewAction(moves []string) *Action {
	return &Action{next: nil, moves: moves}
}

func (a *Action) Execute(p models.Puzzle) Block {
	for _, move := range a.moves {
		p.Turn(move)
	}
	return a.next
}

func (a *Action) NextSetter() Setter {
	return func(block Block) {
		a.next = block
	}
}
