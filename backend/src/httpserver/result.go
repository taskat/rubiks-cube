package httpserver

import (
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/models"
)

type Result[T models.Puzzle] struct {
	State  T                   `json:"state,omitempty"`
	Errors []eh.Message        `json:"errors,omitempty"`
	Turns  map[string][]string `json:"turns,omitempty"`
	Steps  []string            `json:"steps,omitempty"`
}

func NewResult[T models.Puzzle](state T, errors []eh.Message, turns map[string][]string, steps []string) Result[T] {
	return Result[T]{State: state, Errors: errors, Turns: turns, Steps: steps}
}
