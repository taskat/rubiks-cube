package cube

import "github.com/taskat/rubiks-cube/src/color"

type edgePiece struct {
	location     edgeLocation
	colors       [2]color.Color
	goalLocation edgeLocation
}

func newEdgePiece(c *Cube, coords []sideCoord) edgePiece {
	location := newEdgeLocation(coords[0].side, coords[1].side)
	colors := [2]color.Color{c.sides[coords[0].side][coords[0].Row][coords[0].Col],
		c.sides[coords[1].side][coords[1].Row][coords[1].Col]}
	goalLocation := newEdgeLocation(c.getGoalSide(colors[0]), c.getGoalSide(colors[1]))
	return edgePiece{location, colors, goalLocation}
}

func (e edgePiece) movesToGoal() int {
	return e.location.distance(e.goalLocation) + e.costToSolveFlip()
}

func (e edgePiece) costToSolveFlip() int {
	if e.location.same(e.goalLocation) {
		return e.location.flipCost(e.goalLocation)
	}
	if !e.location.hasSameSide(e.goalLocation) {
		e.location = e.location.getMiddleLocation(e.goalLocation)
	}
	if e.isFlipped() {
		return 3
	}
	return 0
}

func (e edgePiece) isFlipped() bool {
	sameSide := e.location.getSameSide(e.goalLocation)
	return (sameSide == e.location[0] && sameSide == e.goalLocation[0]) ||
		(sameSide == e.location[1] && sameSide == e.goalLocation[1])
}

var edgeCoords = [][]sideCoord{
	{newSideCoord(Up, 0, 1), newSideCoord(Back, 0, 1)},
	{newSideCoord(Up, 1, 2), newSideCoord(Right, 0, 1)},
	{newSideCoord(Up, 2, 1), newSideCoord(Front, 0, 1)},
	{newSideCoord(Up, 1, 0), newSideCoord(Left, 0, 1)},
	{newSideCoord(Front, 1, 2), newSideCoord(Right, 1, 0)},
	{newSideCoord(Right, 1, 2), newSideCoord(Back, 1, 0)},
	{newSideCoord(Back, 1, 2), newSideCoord(Left, 1, 0)},
	{newSideCoord(Left, 1, 2), newSideCoord(Front, 1, 0)},
	{newSideCoord(Down, 0, 1), newSideCoord(Front, 2, 1)},
	{newSideCoord(Down, 1, 2), newSideCoord(Right, 2, 1)},
	{newSideCoord(Down, 2, 1), newSideCoord(Back, 2, 1)},
	{newSideCoord(Down, 1, 0), newSideCoord(Left, 2, 1)},
}
