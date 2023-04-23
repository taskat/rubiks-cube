package cube

import (
	"fmt"
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/coord"
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

func NewWithSides(sides map[CubeSide]Side) *Cube {
	return &Cube{
		sides: sides,
		size:  len(sides[Front]),
	}
}

func (c *Cube) getCorners() map[sideCoord]color.Color {
	corners := make(map[sideCoord]color.Color, 24)
	cornerCoords := []coord.Coord{
		{Row: 0, Col: 0},
		{Row: 0, Col: 2},
		{Row: 2, Col: 0},
		{Row: 2, Col: 2},
	}
	for sideName, side := range c.sides {
		for _, cornerCoord := range cornerCoords {
			corners[sideCoord{sideName, cornerCoord}] = side[cornerCoord.Row][cornerCoord.Col]
		}
	}
	return corners
}

func (c *Cube) getEdgePieces() []edgePiece {
	edgePieces := make([]edgePiece, 0, 12)
	for _, edge := range edgeCoords {
		edgePieces = append(edgePieces, newEdgePiece(c, edge))
	}
	return edgePieces
}

func (c *Cube) getGoalSide(color color.Color) CubeSide {
	for sideName, side := range c.sides {
		if side[1][1] == color {
			return sideName
		}
	}
	panic(fmt.Sprintf("No side with color %s", color.String()))
}

func (c *Cube) GetValidator() models.Validator {
	return newValidator(c)
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
