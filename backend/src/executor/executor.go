package executor

import (
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/models"
)

type Executor struct {
	eh *eh.Errorhandler
}

func NewExecutor(eh *eh.Errorhandler) *Executor {
	return &Executor{eh: eh}
}

func (e *Executor) Execute(state models.Puzzle, algo models.Algorithm) []string {
	return nil
}
