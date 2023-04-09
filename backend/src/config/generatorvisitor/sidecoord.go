package generatorvisitor

import (
	"github.com/taskat/rubiks-cube/src/models/coord"
	"github.com/taskat/rubiks-cube/src/models/cube"
)

type sideCoord struct {
	side cube.CubeSide
	coord.Coord
}

func newSideCoord(side cube.CubeSide, row, col int) sideCoord {
	return sideCoord{side, coord.NewCoord(row, col)}
}
