package parameters

type Coord struct {
	Side string
	Row  int
	Col  int
}

func NewCoord(side string, row int, col int) Coord {
	return Coord{side, row, col}
}
