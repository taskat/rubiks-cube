package generatorvisitor

import (
	"github.com/taskat/rubiks-cube/src/color"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/cube"
)

type beginnerStateVisitor struct {
}

func newBeginnerStateVisitor() *beginnerStateVisitor {
	return &beginnerStateVisitor{}
}

func (v *beginnerStateVisitor) visitBeginnerState(ctx *cp.BeginnerStateContext) models.Puzzle {
	sides := make(map[cube.CubeSide]cube.Side, 6)
	for _, side := range ctx.AllSide() {
		cubeSide, sideState := v.visitSide(side.(*cp.SideContext))
		sides[cubeSide] = sideState
	}
	return cube.NewFromBeginnerState(sides)
}

func (v *beginnerStateVisitor) visitColor(ctx *cp.ColorContext) color.Color {
	return color.Color(ctx.GetText())
}

func (v *beginnerStateVisitor) visitSide(ctx *cp.SideContext) (cube.CubeSide, cube.Side) {
	cubeSide := v.visitSideDef(ctx.SideDef().(*cp.SideDefContext))
	side := make(cube.Side, len(ctx.AllRow()))
	for i, row := range ctx.AllRow() {
		side[i] = v.visitRow(row.(*cp.RowContext))
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

func (v *beginnerStateVisitor) visitRow(ctx *cp.RowContext) []color.Color {
	row := make([]color.Color, len(ctx.AllColor()))
	for i, color := range ctx.AllColor() {
		row[i] = v.visitColor(color.(*cp.ColorContext))
	}
	return row
}
