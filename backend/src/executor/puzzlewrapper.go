package executor

import "github.com/taskat/rubiks-cube/src/models"

type puzzleWrapper struct {
	models.Puzzle
	turns []string
}

func NewPuzzleWrapper(puzzle models.Puzzle) *puzzleWrapper {
	return &puzzleWrapper{Puzzle: puzzle, turns: []string{}}
}

func (p *puzzleWrapper) Turn(name string) {
	p.turns = append(p.turns, name)
	p.Puzzle.Turn(name)
}
