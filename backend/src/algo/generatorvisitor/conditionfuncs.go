package generatorvisitor

import (
	"fmt"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/models/parameters"
	"github.com/taskat/rubiks-cube/src/utils"
)

func all(builder conditionBuilderFunc, list parameters.List[parameters.Parameter]) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		for _, param := range list {
			if !builder(param)(p) {
				return false
			}
		}
		return true
	}
}

func any(builder conditionBuilderFunc, list parameters.List[parameters.Parameter]) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		for _, param := range list {
			if builder(param)(p) {
				return true
			}
		}
		return false
	}
}

func at(piece parameters.Piece, pos parameters.Position) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		colors := make([]color.Color, 0, len(pos.Sides))
		for _, sideName := range pos.Sides {
			colors = append(colors, p.GetColor(parameters.NewCoord(sideName, 1, 1)))
		}
		pieceCoords := p.GetPieceCoords(piece)
		posCoords := p.GetPosCoords(pos)
		for _, coord := range pieceCoords {
			if !utils.Contains(posCoords, coord) {
				return false
			}
		}
		return true
	}
}

func colorListMatch(expectedColor color.Color, coords parameters.List[parameters.Coord]) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		fmt.Println("color list match", coords)
		for _, coord := range coords {
			if !p.GetColor(coord).Equals(expectedColor) {
				return false
			}
		}
		return true
	}
}

func like(piece parameters.Piece, pos parameters.Position) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		pieceCoords := p.GetPieceCoords(piece)
		posCoords := p.GetPosCoords(pos)
		for i, coord := range pieceCoords {
			if coord != posCoords[i] {
				return false
			}
		}
		return true
	}
}

func none(builder conditionBuilderFunc, list parameters.List[parameters.Parameter]) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		for _, param := range list {
			if builder(param)(p) {
				return false
			}
		}
		return true
	}
}

func singleColorMatch(expectedColor color.Color, coord parameters.Coord) algorithm.ConditionFunc {
	return func(p models.Puzzle) bool {
		fmt.Println("single color match")
		fmt.Println(p.GetColor(coord))
		return p.GetColor(coord).Equals(expectedColor)
	}
}
