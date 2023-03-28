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
