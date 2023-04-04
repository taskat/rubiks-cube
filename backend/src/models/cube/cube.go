package cube

import (
	"fmt"
	"strings"
)

type Cube struct {
	sides map[CubeSide]Side
	size  int
}

func NewCube(size int) *Cube {
	return &Cube{
		sides: make(map[CubeSide]Side, 6),
		size:  size,
	}
}

func NewFromBeginnerState(sides map[CubeSide]Side) *Cube {
	return &Cube{
		sides: sides,
		size:  len(sides[Front]),
	}
}

func (c *Cube) String() string {
	s := "cube(" + fmt.Sprintf("%d", c.size) + "):\n"
	for sideName, side := range c.sides {
		sideString := side.String()
		lines := strings.Split(sideString, "\n")
		for i, line := range lines {
			lines[i] = "    " + line
		}
		sideString = strings.Join(lines, "\n")
		s += fmt.Sprintf("  %s:\n%s\n", sideName.String(), sideString)
	}
	return s
}

func (c *Cube) MarshalJSON() ([]byte, error) {
	return NewCubeJSON(c).marshal()
}
