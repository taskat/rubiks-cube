package generatorvisitor

import (
	"github.com/taskat/rubiks-cube/src/color"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/cube"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type beginnerStateVisitor struct {
	size            int
	sideConstructor func(string) parameters.Side
}

func newBeginnerStateVisitor(size int, sideConstructor func(string) parameters.Side) *beginnerStateVisitor {
	return &beginnerStateVisitor{size: size, sideConstructor: sideConstructor}
}

func (v *beginnerStateVisitor) visitBeginnerState(ctx *cp.BeginnerStateContext) models.Puzzle {
	sides := make(map[parameters.Side]cube.Side, 6)
	for _, side := range ctx.AllSide() {
		cubeSide, sideState := v.visitSide(side.(*cp.SideContext))
		sides[cubeSide] = sideState
	}
	return cube.NewWithSides(sides)
}

func (v *beginnerStateVisitor) visitColor(ctx *cp.ColorContext) color.Color {
	return color.Color(ctx.GetText())
}

func (v *beginnerStateVisitor) visitSide(ctx *cp.SideContext) (parameters.Side, cube.Side) {
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

func (v *beginnerStateVisitor) visitSideDef(ctx *cp.SideDefContext) parameters.Side {
	return v.sideConstructor(ctx.GetText())
}
