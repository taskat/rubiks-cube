package cube

import (
	"strings"

	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type orientation struct {
	c *Cube
}

func newOrientation() *orientation {
	sides := make(map[parameters.Side]Side, 6)
	sides[Up] = [][]color.Color{{"w"}}
	sides[Down] = [][]color.Color{{"y"}}
	sides[Front] = [][]color.Color{{"b"}}
	sides[Back] = [][]color.Color{{"g"}}
	sides[Left] = [][]color.Color{{"r"}}
	sides[Right] = [][]color.Color{{"o"}}
	c := NewWithSides(sides, 1)
	return &orientation{c: c}
}

func (o *orientation) getColor(side parameters.Side) color.Color {
	return o.c.sides[newCubeSide(side)][0][0]
}

func (o *orientation) turn(move string) {
	if strings.HasPrefix(move, "x") || strings.HasPrefix(move, "y") || strings.HasPrefix(move, "z") {
		o.c.turn(move)
	}
}
