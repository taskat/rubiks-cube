package executor

import (
	"fmt"

	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
)

type Executor struct {
	eh *eh.Errorhandler
}

func NewExecutor(eh *eh.Errorhandler) *Executor {
	return &Executor{eh: eh}
}

func (e *Executor) Execute(state models.Puzzle, algo algorithm.Algorithm) (map[string][]string, []string) {
	puzzleWrapper := newPuzzleWrapper(state)
	block := algo.Execute(puzzleWrapper)
	for !block.Finished() {
		block = block.Execute(puzzleWrapper)
		fmt.Println(block)
	}
	return puzzleWrapper.turns, puzzleWrapper.steps
}
