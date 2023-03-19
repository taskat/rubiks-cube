package cube

type CubeSide int

const (
	Front CubeSide = iota
	Back
	Left
	Right
	Up
	Down
)

func (c CubeSide) String() string {
	s := ""
	switch c {
	case Front:
		s = "Front"
	case Back:
		s = "Back"
	case Left:
		s = "Left"
	case Right:
		s = "Right"
	case Up:
		s = "Up"
	case Down:
		s = "Down"
	default:
		panic("Invalid CubeSide")
	}
	return s
}
