package cube

import (
	"fmt"
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/parameters"
	"github.com/taskat/rubiks-cube/src/utils"
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
	return &c
}

func (c *Cube) Clone() models.Puzzle {
	sides := make(map[CubeSide]Side, 6)
	for sideName, side := range c.sides {
		sides[sideName] = side.clone()
	}
	moves := make(map[string]move, len(c.moves))
	for name, move := range c.moves {
		moves[name] = move
	}
	return &Cube{sides: sides, size: c.size, moves: moves}
}

func (c *Cube) generateMoves() {
	mg := newMoveGenerator(c.size)
	c.moves = mg.generateAllMoves()
}

func (c *Cube) GetColor(coord parameters.Coord) color.Color {
	return c.sides[CubeSide(coord.Side)][coord.Row][coord.Col]
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
	for _, corner := range getCornerCoords() {
		cornerPieces = append(cornerPieces, newCornerPiece(c, corner))
	}
	return cornerPieces
}

func (c *Cube) getEdgePieces() []edgePiece {
	edgePieces := make([]edgePiece, 0, 12)
	for _, edge := range getEdgeCoords() {
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

func (c *Cube) GetPieceCoords(piece parameters.Piece) []parameters.Coord {
	possibleCoords := c.getPossibleCoords(len(piece.Sides))
	for _, coords := range possibleCoords {
		if c.pieceAtCoords(piece, coords) {
			return c.sortCoordsByPiece(coords, piece)
		}
	}
	panic("No coords found")
}

func (c *Cube) GetPosCoords(pos parameters.Position) []parameters.Coord {
	possibleCoords := c.getPossibleCoords(len(pos.Sides))
	for _, coords := range possibleCoords {
		if posAtCoords(pos, coords) {
			return sortCoordsByPos(coords, pos)
		}
	}
	panic("No coords found")
}

func (*Cube) getPossibleCoords(length int) [][]parameters.Coord {
	switch length {
	case 3:
		return getCornerCoords()
	case 2:
		return getEdgeCoords()
	}
	panic("Invalid length")
}

func posAtCoords(pos parameters.Position, coords []parameters.Coord) bool {
	for _, coord := range coords {
		if !utils.Contains(pos.Sides, coord.Side) {
			return false
		}
	}
	return true
}

func (c *Cube) pieceAtCoords(piece parameters.Piece, coords []parameters.Coord) bool {
	colors := make([]color.Color, 0, len(piece.Sides))
	for _, sideName := range piece.Sides {
		colors = append(colors, c.GetColor(parameters.NewCoord(sideName, 1, 1)))
	}
	for _, coord := range coords {
		color := c.GetColor(coord)
		if !utils.Contains(colors, color) {
			return false
		}
	}
	return true
}

func (c *Cube) GetValidator() models.Validator {
	return newValidator(c)
}

func sortCoordsByPos(coords []parameters.Coord, pos parameters.Position) []parameters.Coord {
	sorted := make([]parameters.Coord, len(coords))
	for i, side := range pos.Sides {
		for _, coord := range coords {
			if coord.Side == side {
				sorted[i] = coord
				break
			}
		}
	}
	return sorted
}

func (c *Cube) sortCoordsByPiece(coords []parameters.Coord, piece parameters.Piece) []parameters.Coord {
	sorted := make([]parameters.Coord, len(coords))
	for i, side := range piece.Sides {
		for _, coord := range coords {
			if c.GetColor(coord).Equals(c.GetColor(parameters.NewCoord(side, 1, 1))) {
				sorted[i] = coord
				break
			}
		}
	}
	return sorted
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
				currentColor = c.GetColor(parameters.NewCoordFromString(coord.side.String(), coord.Row, coord.Col))
			}
			nextCoord := cycle[(i+move.steps)%len(cycle)]
			nextColor := c.GetColor(parameters.NewCoordFromString(nextCoord.side.String(), nextCoord.Row, nextCoord.Col))
			moved[nextCoord] = nextColor
			c.sides[nextCoord.side][nextCoord.Row][nextCoord.Col] = currentColor
		}
	}
}

func (c *Cube) MarshalJSON() ([]byte, error) {
	return NewCubeJSON(c).marshal()
}
