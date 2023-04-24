package cube

import "github.com/taskat/rubiks-cube/src/color"

type cornerPiece struct {
	location     cornerLocation
	colors       [3]color.Color
	goalLocation cornerLocation
}

func newCornerPiece(c *Cube, coords []sideCoord) cornerPiece {
	location := newCornerLocation(coords[0].side, coords[1].side, coords[2].side)
	colors := [3]color.Color{}
	for i, coord := range coords {
		colors[i] = c.sides[coord.side][coord.Row][coord.Col]
	}
	goalSides := [3]CubeSide{}
	for i, color := range colors {
		goalSides[i] = c.getGoalSide(color)
	}
	goalLocation := newCornerLocation(goalSides[0], goalSides[1], goalSides[2])
	return cornerPiece{location, colors, goalLocation}
}

var cornerCoords = [][]sideCoord{
	{newSideCoord(Up, 0, 0), newSideCoord(Left, 0, 0), newSideCoord(Back, 0, 2)},
	{newSideCoord(Up, 0, 2), newSideCoord(Back, 0, 0), newSideCoord(Right, 0, 2)},
	{newSideCoord(Up, 2, 2), newSideCoord(Right, 0, 0), newSideCoord(Front, 0, 2)},
	{newSideCoord(Up, 2, 0), newSideCoord(Front, 0, 0), newSideCoord(Left, 0, 2)},
	{newSideCoord(Down, 0, 0), newSideCoord(Front, 2, 0), newSideCoord(Left, 2, 2)},
	{newSideCoord(Down, 0, 2), newSideCoord(Front, 2, 2), newSideCoord(Right, 2, 0)},
	{newSideCoord(Down, 2, 2), newSideCoord(Back, 2, 0), newSideCoord(Right, 2, 2)},
	{newSideCoord(Down, 2, 0), newSideCoord(Back, 2, 2), newSideCoord(Left, 2, 0)},
}
