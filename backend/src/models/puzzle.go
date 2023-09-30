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
	GetPieceCoords(piece parameters.Piece) []parameters.Coord
	GetPosCoords(pos parameters.Position) []parameters.Coord
}

type Validator interface {
	Validate() []string
}
