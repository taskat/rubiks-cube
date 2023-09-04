package cube

import (
	"fmt"
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
)

type Cube struct {
	sides map[CubeSide]Side
	size  int
	moves map[string]move
}

func NewCube(size int) *Cube {
	c := Cube{
		sides: make(map[CubeSide]Side, 6),
		size:  size,
	}
	c.generateMoves()
	return &c
}

func NewWithSides(sides map[CubeSide]Side) *Cube {
	c := Cube{
		sides: sides,
		size:  len(sides[Front]),
	}
	c.generateMoves()
	c.turn("S2")
	c.turn("E2")
	c.turn("M2")
	return &c
}

func (c *Cube) generateMoves() {
	mg := newMoveGenerator(c.size)
	c.moves = mg.generateAllMoves()
}

func (c *Cube) getColor(coord sideCoord) color.Color {
	return c.sides[coord.side][coord.Row][coord.Col]
}

func (c *Cube) GetConstraint() models.Constraint {
	turns := make([]string, 0, len(c.moves))
	for name := range c.moves {
		turns = append(turns, name)
	}
	colors := []string{"white", "red", "orange", "yellow", "green", "blue"}
	return *models.NewConstraint(turns, AllSides(), colors)
}

func (c *Cube) getCornerPieces() []cornerPiece {
	cornerPieces := make([]cornerPiece, 0, 8)
	for _, corner := range cornerCoords {
		cornerPieces = append(cornerPieces, newCornerPiece(c, corner))
	}
	return cornerPieces
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

func (c *Cube) Turn(name string) {
	_, ok := c.moves[name]
	if !ok {
		panic(fmt.Sprintf("Undefined move %s", name))
	}
	c.turn(name)
}

func (c *Cube) turn(name string) {
	move := c.moves[name]
	for _, cycle := range move.cycles {
		moved := make(map[sideCoord]color.Color)
		for i, coord := range cycle {
			var currentColor color.Color
			if color, ok := moved[coord]; ok {
				currentColor = color
			} else {
				currentColor = c.getColor(coord)
			}
			nextCoord := cycle[(i+move.steps)%len(cycle)]
			nextColor := c.getColor(nextCoord)
			moved[nextCoord] = nextColor
			c.sides[nextCoord.side][nextCoord.Row][nextCoord.Col] = currentColor
		}
	}
}

func (c *Cube) MarshalJSON() ([]byte, error) {
	return NewCubeJSON(c).marshal()
}
