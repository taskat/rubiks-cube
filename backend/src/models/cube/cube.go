package cube

import (
	"fmt"
	"strings"
)

type Cube struct {
	sides map[CubeSide]side
	size  int
}

func NewCube(size int) *Cube {
	return &Cube{
		sides: make(map[CubeSide]side, 6),
		size:  size,
	}
}

func (c *Cube) String() string {
	s := "cube(" + fmt.Sprintf("%d", c.size) + "): "
	for sideName, side := range c.sides {
		sideString := side.String()
		lines := strings.Split(sideString, "\n")
		for i, line := range lines {
			lines[i] = "  " + line
		}
		sideString = strings.Join(lines, "\n")
		s += fmt.Sprintf("%s:\n%s, ", sideName.String(), side.String())
	}
	return s
}
