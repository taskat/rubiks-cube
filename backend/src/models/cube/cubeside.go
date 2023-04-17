package cube

import "github.com/taskat/rubiks-cube/src/models/coord"

type CubeSide string

const (
	Front CubeSide = "Front"
	Back  CubeSide = "Back"
	Left  CubeSide = "Left"
	Right CubeSide = "Right"
	Up    CubeSide = "Up"
	Down  CubeSide = "Down"
)

func (c CubeSide) String() string {
	return string(c)
}

type sideCoord struct {
	side CubeSide
	coord.Coord
}
