package cube

import "github.com/taskat/rubiks-cube/src/models/coord"

type sideCoord struct {
	side CubeSide
	coord.Coord
}

func newSideCoord(side CubeSide, row, col int) sideCoord {
	return sideCoord{
		side:  side,
		Coord: coord.NewCoord(row, col),
	}
}
