package generatorvisitor

import (
	"fmt"
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type parameterVisitor struct{}

func newParameterVisitor() *parameterVisitor {
	return &parameterVisitor{}
}

func (v *parameterVisitor) visitCoord(ctx *ap.CoordContext) parameters.Coord {
	side := ctx.Side().GetText()
	row, _ := strconv.Atoi(ctx.NUMBER(0).GetText())
	col, _ := strconv.Atoi(ctx.NUMBER(1).GetText())
	return parameters.NewCoord(side, row, col)
}

func (v *parameterVisitor) visitParameter(ctx *ap.ParameterContext) parameters.Parameter {
	if ctx.Coord() != nil {
		return v.visitCoord(ctx.Coord().(*ap.CoordContext))
	}
	fmt.Println("Unknown parameter type")
	return nil
}
