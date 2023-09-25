package models

import (
	"encoding/json"
	"fmt"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type Puzzle interface {
	json.Marshaler
	fmt.Stringer
	Clone() Puzzle
	GetValidator() Validator
	GetConstraint() Constraint
	Turn(name string)
	GetColor(coord parameters.Coord) color.Color
	PieceLike(piece parameters.Piece, pos parameters.Position) bool
	PieceAt(piece parameters.Piece, pos parameters.Position) bool
}

type Validator interface {
	Validate() []string
}
