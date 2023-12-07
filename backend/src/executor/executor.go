package executor

import (
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/models/panics"
	"github.com/taskat/rubiks-cube/src/utils"
)

type Executor struct {
	eh *eh.Errorhandler
}

func NewExecutor(eh *eh.Errorhandler) *Executor {
	return &Executor{eh: eh}
}

func (e *Executor) Execute(state models.Puzzle, algo algorithm.Algorithm) (result utils.Result[ExecutionResult]) {
	defer e.handleMaxRuns(&result)
	puzzleWrapper := newPuzzleWrapper(state)
	block := algo.Execute(puzzleWrapper)
	for !block.Finished() {
		block = block.Execute(puzzleWrapper)
	}
	executionResult := *NewExecutionResult(puzzleWrapper.turns, puzzleWrapper.steps)
	return utils.Ok(executionResult)
}

func (e *Executor) handleMaxRuns(result *utils.Result[ExecutionResult]) {
	if r := recover(); r != nil {
		if maxRunReached, ok := r.(*panics.MaxRunReached); ok {
			*result = utils.Err[ExecutionResult](maxRunReached)
		} else {
			panic(r)
		}
	}
}
