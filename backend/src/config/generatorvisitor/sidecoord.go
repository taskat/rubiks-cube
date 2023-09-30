package generatorvisitor

import (
	"github.com/taskat/rubiks-cube/src/models/coord"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type sideCoord struct {
	side parameters.Side
	coord.Coord
}

func newSideCoord(side parameters.Side, row, col int) sideCoord {
	return sideCoord{side, coord.NewCoord(row, col)}
}
