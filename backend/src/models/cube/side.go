package cube

import (
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
)

type Side [][]color.Color

func (s Side) String() string {
	lines := make([]string, len(s))
	for i, row := range s {
		cells := make([]string, len(row))
		for j, color := range row {
			cells[j] = color.String()
		}
		lines[i] = strings.Join(cells, " ")
	}
	return strings.Join(lines, "\n")
}

func (s Side) clone() Side {
	side := make(Side, len(s))
	for i, row := range s {
		side[i] = make([]color.Color, len(row))
		copy(side[i], row)
	}
	return side
}
