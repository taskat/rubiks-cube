package generatorvisitor

import (
	"github.com/taskat/rubiks-cube/src/color"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/cube"
)

type beginnerStateVisitor struct {
	size int
}

func newBeginnerStateVisitor(size int) *beginnerStateVisitor {
	return &beginnerStateVisitor{size: size}
}

func (v *beginnerStateVisitor) visitBeginnerState(ctx *cp.BeginnerStateContext) models.Puzzle {
	sides := make(map[cube.CubeSide]cube.Side, 6)
	for _, side := range ctx.AllSide() {
		cubeSide, sideState := v.visitSide(side.(*cp.SideContext))
		sides[cubeSide] = sideState
	}
	return cube.NewWithSides(sides)
}

func (v *beginnerStateVisitor) visitColor(ctx *cp.ColorContext) color.Color {
	return color.Color(ctx.GetText())
}

func (v *beginnerStateVisitor) visitSide(ctx *cp.SideContext) (cube.CubeSide, cube.Side) {
	cubeSide := v.visitSideDef(ctx.SideDef().(*cp.SideDefContext))
	side := make(cube.Side, v.size)
	for i := range side {
		side[i] = make([]color.Color, v.size)
		for j := range side[i] {
			side[i][j] = color.Color("-")
		}
	}
	for i, color := range ctx.AllColor() {
		side[i/v.size][i%v.size] = v.visitColor(color.(*cp.ColorContext))
	}
	return cubeSide, side
}

func (v *beginnerStateVisitor) visitSideDef(ctx *cp.SideDefContext) cube.CubeSide {
	switch ctx.GetText() {
	case "front":
		return cube.Front
	case "back":
		return cube.Back
	case "up":
		return cube.Up
	case "down":
		return cube.Down
	case "left":
		return cube.Left
	case "right":
		return cube.Right
	default:
		panic("invalid side")
	}
}
