package executor

import "github.com/taskat/rubiks-cube/src/models"

type puzzleWrapper struct {
	models.Puzzle
	turns map[string][]string
	steps []string
}

func newPuzzleWrapper(puzzle models.Puzzle) *puzzleWrapper {
	return &puzzleWrapper{Puzzle: puzzle, turns: make(map[string][]string), steps: make([]string, 0)}
}

func (p *puzzleWrapper) currentStep() string {
	return p.steps[len(p.steps)-1]
}

func (p *puzzleWrapper) Turn(name string) {
	p.turns[p.currentStep()] = append(p.turns[p.currentStep()], name)
	p.Puzzle.Turn(name)
}

func (p *puzzleWrapper) StartStep(name string) {
	p.steps = append(p.steps, name)
	p.turns[name] = make([]string, 0)
}
