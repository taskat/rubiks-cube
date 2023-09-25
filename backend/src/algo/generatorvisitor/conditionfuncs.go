package generatorvisitor

import (
	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

func at(piece parameters.Piece, pos parameters.Position) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		return p.PieceAt(piece, pos)
	}
}

func colorListMatch(expectedColor color.Color, coords parameters.List[parameters.Coord]) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		for _, coord := range coords {
			if p.GetColor(coord) != expectedColor {
				return false
			}
		}
		return true
	}
}

func like(piece parameters.Piece, pos parameters.Position) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		return p.PieceLike(piece, pos)
	}
}

func singleColorMatch(expectedColor color.Color, coord parameters.Coord) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		return p.GetColor(coord)[0] == expectedColor[0]
	}
}
