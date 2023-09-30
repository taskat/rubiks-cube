package cube

import (
	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type cornerPiece struct {
	location     cornerLocation
	colors       [3]color.Color
	goalLocation cornerLocation
}

func newCornerPiece(c *Cube, coords []parameters.Coord) cornerPiece {
	location := newCornerLocation(CubeSide(coords[0].Side), CubeSide(coords[1].Side), CubeSide(coords[2].Side))
	colors := [3]color.Color{}
	for i, coord := range coords {
		colors[i] = c.sides[CubeSide(coord.Side)][coord.Row][coord.Col]
	}
	goalSides := [3]CubeSide{}
	for i, color := range colors {
		goalSides[i] = c.getGoalSide(color)
	}
	goalLocation := newCornerLocation(goalSides[0], goalSides[1], goalSides[2])
	return cornerPiece{location, colors, goalLocation}
}

func (c cornerPiece) getTwist(isGood func(color.Color) bool) int {
	for i, color := range c.colors {
		if isGood(color) {
			return i
		}
	}
	panic("No good side found")
}

func getCornerCoords() [][]parameters.Coord {
	coords := make([][]parameters.Coord, 8)
	coords[0] = []parameters.Coord{
		parameters.NewCoord("Up", 0, 0),
		parameters.NewCoord("Left", 0, 0),
		parameters.NewCoord("Back", 0, 2),
	}
	coords[1] = []parameters.Coord{
		parameters.NewCoord("Up", 0, 2),
		parameters.NewCoord("Back", 0, 0),
		parameters.NewCoord("Right", 0, 2),
	}
	coords[2] = []parameters.Coord{
		parameters.NewCoord("Up", 2, 2),
		parameters.NewCoord("Right", 0, 0),
		parameters.NewCoord("Front", 0, 2),
	}
	coords[3] = []parameters.Coord{
		parameters.NewCoord("Up", 2, 0),
		parameters.NewCoord("Front", 0, 0),
		parameters.NewCoord("Left", 0, 2),
	}
	coords[4] = []parameters.Coord{
		parameters.NewCoord("Down", 0, 0),
		parameters.NewCoord("Left", 2, 2),
		parameters.NewCoord("Front", 2, 0),
	}
	coords[5] = []parameters.Coord{
		parameters.NewCoord("Down", 0, 2),
		parameters.NewCoord("Front", 2, 2),
		parameters.NewCoord("Right", 2, 0),
	}
	coords[6] = []parameters.Coord{
		parameters.NewCoord("Down", 2, 2),
		parameters.NewCoord("Right", 2, 2),
		parameters.NewCoord("Back", 2, 0),
	}
	coords[7] = []parameters.Coord{
		parameters.NewCoord("Down", 2, 0),
		parameters.NewCoord("Back", 2, 2),
		parameters.NewCoord("Left", 2, 0),
	}
	return coords
}
