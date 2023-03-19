package cube

import (
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
)

type side [][]color.Color

func (s side) String() string {
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
