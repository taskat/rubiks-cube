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

func getEdgeCoords(size int) [][]parameters.Coord {
	perEdge := size - 2
	coords := make([][]parameters.Coord, 12*perEdge)
	inverse := func(i int) int {
		return size - i - 1
	}
	max := size - 1
	for i := 1; i <= perEdge; i++ {
		ub := []parameters.Coord{
			parameters.NewCoord(cubeSide("Up"), 0, i),
			parameters.NewCoord(cubeSide("Back"), 0, inverse(i)),
		}
		coords = append(coords, ub)
		ur := []parameters.Coord{
			parameters.NewCoord(cubeSide("Up"), i, max),
			parameters.NewCoord(cubeSide("Right"), 0, inverse(i)),
		}
		coords = append(coords, ur)
		uf := []parameters.Coord{
			parameters.NewCoord(cubeSide("Up"), max, inverse(i)),
			parameters.NewCoord(cubeSide("Front"), 0, inverse(i)),
		}
		coords = append(coords, uf)
		ul := []parameters.Coord{
			parameters.NewCoord(cubeSide("Up"), inverse(i), 0),
			parameters.NewCoord(cubeSide("Left"), 0, inverse(i)),
		}
		coords = append(coords, ul)
		df := []parameters.Coord{
			parameters.NewCoord(cubeSide("Down"), 0, i),
			parameters.NewCoord(cubeSide("Front"), max, i),
		}
		coords = append(coords, df)
		dr := []parameters.Coord{
			parameters.NewCoord(cubeSide("Down"), i, max),
			parameters.NewCoord(cubeSide("Right"), max, i),
		}
		coords = append(coords, dr)
		db := []parameters.Coord{
			parameters.NewCoord(cubeSide("Down"), max, inverse(i)),
			parameters.NewCoord(cubeSide("Back"), max, i),
		}
		coords = append(coords, db)
		dl := []parameters.Coord{
			parameters.NewCoord(cubeSide("Down"), inverse(i), 0),
			parameters.NewCoord(cubeSide("Left"), max, i),
		}
		coords = append(coords, dl)
		lb := []parameters.Coord{
			parameters.NewCoord(cubeSide("Left"), i, 0),
			parameters.NewCoord(cubeSide("Back"), i, max),
		}
		coords = append(coords, lb)
		lf := []parameters.Coord{
			parameters.NewCoord(cubeSide("Left"), i, max),
			parameters.NewCoord(cubeSide("Front"), i, 0),
		}
		coords = append(coords, lf)
		rb := []parameters.Coord{
			parameters.NewCoord(cubeSide("Right"), i, max),
			parameters.NewCoord(cubeSide("Back"), i, 0),
		}
		coords = append(coords, rb)
		rf := []parameters.Coord{
			parameters.NewCoord(cubeSide("Right"), i, 0),
			parameters.NewCoord(cubeSide("Front"), i, max),
		}
		coords = append(coords, rf)
	}
	return coords
}
