package cube

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

func (c CubeSide) getOpposite() CubeSide {
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

func (c CubeSide) getHash() string {
	return string(c[0])
}

func (c CubeSide) isOpposite(other CubeSide) bool {
	opposite := c.getOpposite()
	return opposite == other
}

func AllSides() []string {
	return []string{Front.String(), Back.String(), Left.String(),
		Right.String(), Up.String(), Down.String()}
}
