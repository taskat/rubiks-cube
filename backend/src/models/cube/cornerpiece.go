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
	location := newCornerLocation(coords[0].Side, coords[1].Side, coords[2].Side)
	colors := [3]color.Color{}
	for i, coord := range coords {
		colors[i] = c.sides[newCubeSide(coord.Side)][coord.Row][coord.Col]
	}
	goalSides := [3]cubeSide{}
	for i, color := range colors {
		goalSides[i] = c.getGoalSide(color)
	}
	goalLocation := newCornerLocationFromCubeSides(goalSides[0], goalSides[1], goalSides[2])
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

func getCornerCoords(size int) [][]parameters.Coord {
	coords := make([][]parameters.Coord, 8)
	max := size - 1
	coords[0] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), 0, 0),
		parameters.NewCoord(cubeSide("Left"), 0, 0),
		parameters.NewCoord(cubeSide("Back"), 0, max),
	}
	coords[1] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), 0, max),
		parameters.NewCoord(cubeSide("Back"), 0, 0),
		parameters.NewCoord(cubeSide("Right"), 0, max),
	}
	coords[2] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), max, max),
		parameters.NewCoord(cubeSide("Right"), 0, 0),
		parameters.NewCoord(cubeSide("Front"), 0, max),
	}
	coords[3] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), max, 0),
		parameters.NewCoord(cubeSide("Front"), 0, 0),
		parameters.NewCoord(cubeSide("Left"), 0, max),
	}
	coords[4] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), 0, 0),
		parameters.NewCoord(cubeSide("Left"), max, max),
		parameters.NewCoord(cubeSide("Front"), max, 0),
	}
	coords[5] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), 0, max),
		parameters.NewCoord(cubeSide("Front"), max, max),
		parameters.NewCoord(cubeSide("Right"), max, 0),
	}
	coords[6] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), max, max),
		parameters.NewCoord(cubeSide("Right"), max, max),
		parameters.NewCoord(cubeSide("Back"), max, 0),
	}
	coords[7] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), max, 0),
		parameters.NewCoord(cubeSide("Back"), max, max),
		parameters.NewCoord(cubeSide("Left"), max, 0),
	}
	return coords
}
