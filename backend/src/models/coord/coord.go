package coord

type Coord struct {
	Row int
	Col int
}

func NewCoord(row, col int) Coord {
	return Coord{row, col}
}
