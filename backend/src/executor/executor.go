package executor

import (
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

func (e *Executor) Execute(state models.Puzzle, algo algorithm.Algorithm) []string {
	return nil
}
