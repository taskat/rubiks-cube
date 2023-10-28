package parameters

type Coord struct {
	Side Side
	Row  int
	Col  int
}

func NewCoord(side Side, row int, col int) Coord {
	return Coord{side, row, col}
}
