package cube

import "github.com/taskat/rubiks-cube/src/models/parameters"

type cubeSide parameters.Side

const (
	Front cubeSide = "Front"
	Back  cubeSide = "Back"
	Left  cubeSide = "Left"
	Right cubeSide = "Right"
	Up    cubeSide = "Up"
	Down  cubeSide = "Down"
)

func (c cubeSide) String() string {
	return string(c)
}

func (c cubeSide) getOpposite() cubeSide {
	switch c {
	case Front:
		return Back
	case Back:
		return Front
	case Left:
		return Right
	case Right:
		return Left
	case Up:
		return Down
	case Down:
		return Up
	}
	panic("Invalid side")
}

func (c cubeSide) getHash() string {
	return string(c[0])
}

func (c cubeSide) isOpposite(other cubeSide) bool {
	opposite := c.getOpposite()
	return opposite == other
}

func AllSides() []string {
	return []string{Front.String(), Back.String(), Left.String(),
		Right.String(), Up.String(), Down.String()}
}
