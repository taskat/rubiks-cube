package cube

import (
	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type edgePiece struct {
	location     edgeLocation
	colors       [2]color.Color
	goalLocation edgeLocation
	c            *Cube
}

func newEdgePiece(c *Cube, coords []parameters.Coord) edgePiece {
	location := newEdgeLocation(coords[0].Side, coords[1].Side)
	colors := [2]color.Color{c.sides[newCubeSide(coords[0].Side)][coords[0].Row][coords[0].Col],
		c.sides[newCubeSide(coords[1].Side)][coords[1].Row][coords[1].Col]}
	goalLocation := newEdgeLocationFromCubeSides(c.getGoalSide(colors[0]), c.getGoalSide(colors[1]))
	return edgePiece{location, colors, goalLocation, c}
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

	if e.location.sides[0] == sameSide {
		return e.c.getGoalSide(e.colors[0]) != sameSide
	}
	return e.c.getGoalSide(e.colors[1]) != sameSide
}

func getEdgeCoords() [][]parameters.Coord {
	coords := make([][]parameters.Coord, 12)
	coords[0] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), 0, 1),
		parameters.NewCoord(cubeSide("Back"), 0, 1),
	}
	coords[1] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), 1, 2),
		parameters.NewCoord(cubeSide("Right"), 0, 1),
	}
	coords[2] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), 2, 1),
		parameters.NewCoord(cubeSide("Front"), 0, 1),
	}
	coords[3] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Up"), 1, 0),
		parameters.NewCoord(cubeSide("Left"), 0, 1),
	}
	coords[4] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), 0, 1),
		parameters.NewCoord(cubeSide("Front"), 2, 1),
	}
	coords[5] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), 1, 2),
		parameters.NewCoord(cubeSide("Right"), 2, 1),
	}
	coords[6] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), 2, 1),
		parameters.NewCoord(cubeSide("Back"), 2, 1),
	}
	coords[7] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Down"), 1, 0),
		parameters.NewCoord(cubeSide("Left"), 2, 1),
	}
	coords[8] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Left"), 1, 0),
		parameters.NewCoord(cubeSide("Back"), 1, 2),
	}
	coords[9] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Left"), 1, 2),
		parameters.NewCoord(cubeSide("Front"), 1, 0),
	}
	coords[10] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Right"), 1, 0),
		parameters.NewCoord(cubeSide("Front"), 1, 2),
	}
	coords[11] = []parameters.Coord{
		parameters.NewCoord(cubeSide("Right"), 1, 2),
		parameters.NewCoord(cubeSide("Back"), 1, 0),
	}
	return coords
}
