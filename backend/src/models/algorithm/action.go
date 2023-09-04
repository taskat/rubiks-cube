package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type action struct {
	puzzle models.Puzzle
	next   Block
	moves  []string
}

func newAction(puzzle models.Puzzle, next Block, moves []string) Block {
	return &action{puzzle: puzzle, next: next, moves: moves}
}

func (a *action) Execute() Block {
	for _, move := range a.moves {
		a.puzzle.Turn(move)
	}
	return a.next
}
