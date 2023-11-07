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
	sides       map[cubeSide]Side
	size        int
	moves       map[string]move
	orientation *orientation
}

func NewCube(size int) *Cube {
	sides := make(map[parameters.Side]Side, 6)
	return NewWithSides(sides, size)
}

func NewWithSides(sides map[parameters.Side]Side, size int) *Cube {
	cubeSides := make(map[cubeSide]Side, 6)
	for sideName, side := range sides {
		cubeSides[newCubeSide(sideName)] = side
	}
	c := Cube{
		sides: cubeSides,
		size:  size,
	}
	if size%2 == 0 {
		c.orientation = newOrientation()
	}
	c.generateMoves()
	return &c
}

func (c *Cube) Clone() models.Puzzle {
	sides := make(map[cubeSide]Side, 6)
	for sideName, side := range c.sides {
		sides[sideName] = side.clone()
	}
	moves := make(map[string]move, len(c.moves))
	for name, move := range c.moves {
		moves[name] = move
	}
	return &Cube{sides: sides, size: c.size, moves: moves, orientation: c.orientation}
}

func (c *Cube) generateMoves() {
	mg := newMoveGenerator(c.size)
	c.moves = mg.generateAllMoves()
}

func (c *Cube) GetColor(coord parameters.Coord) color.Color {
	return c.sides[newCubeSide(coord.Side)][coord.Row][coord.Col]
}

func (c *Cube) GetConstraint() models.Constraint {
	turns := make([]string, 0, len(c.moves))
	for name := range c.moves {
		turns = append(turns, name)
	}
	colors := []string{"white", "red", "orange", "yellow", "green", "blue"}
	return *models.NewConstraint(turns, AllSides(), colors, c.size)
}

func (c *Cube) getCornerPieces() []cornerPiece {
	cornerPieces := make([]cornerPiece, 0, 8)
	for _, corner := range getCornerCoords(c.size) {
		cornerPieces = append(cornerPieces, newCornerPiece(c, corner))
	}
	return cornerPieces
}

func (c *Cube) getMiddleEdgePieces() []edgePiece {
	edgePieces := make([]edgePiece, 0, 12*(c.size-2))
	middle := c.size / 2
	for _, edge := range getEdgeCoords(c.size) {
		if edge[0].Row != middle || edge[0].Col != middle {
			continue
		}
		edgePieces = append(edgePieces, newEdgePiece(c, edge))
	}
	return edgePieces
}

func (c *Cube) getGoalSide(col color.Color) cubeSide {
	if c.size%2 == 0 {
		switch col {
		case color.Color("w"):
			return cubeSide("Up")
		case color.Color("y"):
			return cubeSide("Down")
		case color.Color("r"):
			return cubeSide("Left")
		case color.Color("o"):
			return cubeSide("Right")
		case color.Color("g"):
			return cubeSide("Back")
		case color.Color("b"):
			return cubeSide("Front")
		default:
			panic(fmt.Sprintf("Invalid color %s", col.String()))
		}
	}
	for sideName, side := range c.sides {
		middle := c.size / 2
		if side[middle][middle] == col {
			return sideName
		}
	}
	panic(fmt.Sprintf("No side with color %s", col.String()))
}

func (c *Cube) GetPieceCoords(piece parameters.Piece) []parameters.Coord {
	possibleCoords := c.getPossibleCoords(len(piece.Sides))
	for _, coords := range possibleCoords {
		if c.pieceAtCoords(piece, coords) {
			return c.sortCoordsByPiece(coords, piece)
		}
	}
	fmt.Println(piece)
	fmt.Println(possibleCoords)
	c.pieceAtCoords(piece, possibleCoords[0], true)
	fmt.Println(c.orientation)
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

func (c *Cube) getPossibleCoords(length int) [][]parameters.Coord {
	switch length {
	case 3:
		return getCornerCoords(c.size)
	case 2:
		return getEdgeCoords(c.size)
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

func (c *Cube) getSideColor(side parameters.Side) color.Color {
	if c.orientation == nil {
		middle := c.size / 2
		return c.GetColor(parameters.NewCoord(side, middle, middle))
	}
	return c.orientation.getColor(side)
}

func (c *Cube) pieceAtCoords(piece parameters.Piece, coords []parameters.Coord, log ...bool) bool {
	colors := make([]color.Color, 0, len(piece.Sides))
	for _, sideName := range piece.Sides {
		colors = append(colors, c.getSideColor(sideName))
	}
	if log != nil && log[0] {
		fmt.Println(piece)
		fmt.Println(coords)
		fmt.Println(colors)
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
			if c.GetColor(coord).Equals(c.getSideColor(side)) {
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
		moved := make(map[parameters.Coord]color.Color)
		for i, coord := range cycle {
			var currentColor color.Color
			if color, ok := moved[coord]; ok {
				currentColor = color
			} else {
				currentColor = c.GetColor(coord)
			}
			nextCoord := cycle[(i+move.steps)%len(cycle)]
			nextColor := c.GetColor(nextCoord)
			moved[nextCoord] = nextColor
			c.sides[newCubeSide(nextCoord.Side)][nextCoord.Row][nextCoord.Col] = currentColor
		}
	}
	if c.orientation != nil {
		c.orientation.turn(name)
	}
}

func (c *Cube) MarshalJSON() ([]byte, error) {
	return NewCubeJSON(c).marshal()
}

func (c *Cube) SideConstructor() func(s string) parameters.Side {
	return NewCubeSide
}
