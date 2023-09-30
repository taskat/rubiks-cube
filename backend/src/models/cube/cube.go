package cube

import (
	"fmt"
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/parameters"
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

func (c *Cube) getCornerCoords() [][]parameters.Coord {
	coords := make([][]parameters.Coord, 8)
	coords[0] = []parameters.Coord{
		parameters.NewCoord("Up", 0, 0),
		parameters.NewCoord("Left", 0, 0),
		parameters.NewCoord("Back", 0, 2),
	}
	coords[1] = []parameters.Coord{
		parameters.NewCoord("Up", 0, 2),
		parameters.NewCoord("Back", 0, 0),
		parameters.NewCoord("Right", 0, 2),
	}
	coords[2] = []parameters.Coord{
		parameters.NewCoord("Up", 2, 2),
		parameters.NewCoord("Right", 0, 0),
		parameters.NewCoord("Front", 0, 2),
	}
	coords[3] = []parameters.Coord{
		parameters.NewCoord("Up", 2, 0),
		parameters.NewCoord("Front", 0, 0),
		parameters.NewCoord("Left", 0, 2),
	}
	coords[4] = []parameters.Coord{
		parameters.NewCoord("Down", 0, 0),
		parameters.NewCoord("Left", 2, 2),
		parameters.NewCoord("Front", 2, 0),
	}
	coords[5] = []parameters.Coord{
		parameters.NewCoord("Down", 0, 2),
		parameters.NewCoord("Front", 2, 2),
		parameters.NewCoord("Right", 2, 0),
	}
	coords[6] = []parameters.Coord{
		parameters.NewCoord("Down", 2, 2),
		parameters.NewCoord("Right", 2, 2),
		parameters.NewCoord("Back", 2, 0),
	}
	coords[7] = []parameters.Coord{
		parameters.NewCoord("Down", 2, 0),
		parameters.NewCoord("Back", 2, 2),
		parameters.NewCoord("Left", 2, 0),
	}
	return coords
}

func (c *Cube) getEdgeCoords() [][]parameters.Coord {
	coords := make([][]parameters.Coord, 12)
	coords[0] = []parameters.Coord{
		parameters.NewCoord("Up", 0, 1),
		parameters.NewCoord("Back", 0, 1),
	}
	coords[1] = []parameters.Coord{
		parameters.NewCoord("Up", 1, 2),
		parameters.NewCoord("Right", 0, 1),
	}
	coords[2] = []parameters.Coord{
		parameters.NewCoord("Up", 2, 1),
		parameters.NewCoord("Front", 0, 1),
	}
	coords[3] = []parameters.Coord{
		parameters.NewCoord("Up", 1, 0),
		parameters.NewCoord("Left", 0, 1),
	}
	coords[4] = []parameters.Coord{
		parameters.NewCoord("Down", 0, 1),
		parameters.NewCoord("Front", 2, 1),
	}
	coords[5] = []parameters.Coord{
		parameters.NewCoord("Down", 1, 2),
		parameters.NewCoord("Right", 2, 1),
	}
	coords[6] = []parameters.Coord{
		parameters.NewCoord("Down", 2, 1),
		parameters.NewCoord("Back", 2, 1),
	}
	coords[7] = []parameters.Coord{
		parameters.NewCoord("Down", 1, 0),
		parameters.NewCoord("Left", 2, 1),
	}
	coords[8] = []parameters.Coord{
		parameters.NewCoord("Left", 1, 0),
		parameters.NewCoord("Back", 1, 2),
	}
	coords[9] = []parameters.Coord{
		parameters.NewCoord("Left", 1, 2),
		parameters.NewCoord("Front", 1, 0),
	}
	coords[10] = []parameters.Coord{
		parameters.NewCoord("Right", 1, 0),
		parameters.NewCoord("Front", 1, 2),
	}
	coords[11] = []parameters.Coord{
		parameters.NewCoord("Right", 1, 2),
		parameters.NewCoord("Back", 1, 0),
	}
	return coords
}

func (c *Cube) GetPieceCoords(piece parameters.Piece) []parameters.Coord {
	colors := make([]color.Color, 0, len(piece.Sides))
	for _, sideName := range piece.Sides {
		colors = append(colors, c.sides[CubeSide(sideName)][1][1])
	}
	var possibleCoords [][]parameters.Coord
	switch len(piece.Sides) {
	case 3:
		possibleCoords = c.getCornerCoords()
	case 2:
		possibleCoords = c.getEdgeCoords()
	}
	for _, coords := range possibleCoords {
		if c.checkCoords(coords, colors) {
			return c.sortByPiece(coords, piece)
		}
	}
	panic("No coords found")
}

func (c *Cube) GetPosCoords(pos parameters.Position) []parameters.Coord {
	var possibleCoords [][]parameters.Coord
	switch len(pos.Sides) {
	case 3:
		possibleCoords = c.getCornerCoords()
	case 2:
		possibleCoords = c.getEdgeCoords()
	}
	for _, coords := range possibleCoords {
		if checkPos(coords, pos) {
			return sortByPos(coords, pos)
		}
	}
	panic("No coords found")
}

func checkPos(coords []parameters.Coord, pos parameters.Position) bool {
	for _, coord := range coords {
		if !containsCoord(pos, coord) {
			return false
		}
	}
	return true
}

func containsCoord(pos parameters.Position, coord parameters.Coord) bool {
	for _, posSide := range pos.Sides {
		if posSide == coord.Side {
			return true
		}
	}
	return false
}

func containsCoordCoord(coords []parameters.Coord, elem parameters.Coord) bool {
	for _, coord := range coords {
		if coord == elem {
			return true
		}
	}
	return false
}

func (c *Cube) checkCoords(coords []parameters.Coord, colors []color.Color) bool {
	for _, coord := range coords {
		color := c.GetColor(coord)
		if !contain(colors, color) {
			return false
		}
	}
	return true
}

func contain(arr []color.Color, elem color.Color) bool {
	for _, e := range arr {
		if e[0] == elem[0] {
			return true
		}
	}
	return false
}

func (c *Cube) GetValidator() models.Validator {
	return newValidator(c)
}

func sortByPos(coords []parameters.Coord, pos parameters.Position) []parameters.Coord {
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

func (c *Cube) sortByPiece(coords []parameters.Coord, piece parameters.Piece) []parameters.Coord {
	sorted := make([]parameters.Coord, len(coords))
	for i, side := range piece.Sides {
		for _, coord := range coords {
			if c.GetColor(coord)[0] == c.GetColor(parameters.NewCoord(side, 1, 1))[0] {
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
