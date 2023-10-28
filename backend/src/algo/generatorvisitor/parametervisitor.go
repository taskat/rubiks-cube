package generatorvisitor

import (
	"fmt"
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type parameterVisitor struct {
	functionalEnabled bool
	sideConstructor   func(s string) parameters.Side
}

func newParameterVisitor(functionalEnabled bool, sideConstructor func(s string) parameters.Side) *parameterVisitor {
	return &parameterVisitor{functionalEnabled: functionalEnabled, sideConstructor: sideConstructor}
}

func (v *parameterVisitor) visitCoord(ctx *ap.CoordContext) parameters.Coord {
	side := ctx.Side().GetText()
	row, _ := strconv.Atoi(ctx.NUMBER(0).GetText())
	col, _ := strconv.Atoi(ctx.NUMBER(1).GetText())
	return parameters.NewCoord(v.sideConstructor(side), row, col)
}

func (v *parameterVisitor) visitList(ctx *ap.ListContext) parameters.List[parameters.Parameter] {
	if ctx.Node(0) != nil {
		list := make(parameters.List[parameters.Parameter], len(ctx.AllNode()))
		for i, node := range ctx.AllNode() {
			list[i] = v.visitNode(node.(*ap.NodeContext))
		}
		return list
	}
	if ctx.Piece(0) != nil {
		list := make(parameters.List[parameters.Parameter], len(ctx.AllPiece()))
		for i, piece := range ctx.AllPiece() {
			list[i] = v.visitPiece(piece.(*ap.PieceContext))
		}
		return list
	}
	if ctx.Position(0) != nil {
		list := make(parameters.List[parameters.Parameter], len(ctx.AllPosition()))
		for i, position := range ctx.AllPosition() {
			list[i] = v.visitPosition(position.(*ap.PositionContext))
		}
		return list
	}
	if ctx.Coord(0) != nil {
		list := make(parameters.List[parameters.Parameter], len(ctx.AllCoord()))
		for i, coord := range ctx.AllCoord() {
			list[i] = v.visitCoord(coord.(*ap.CoordContext))
		}
		return list
	}
	panic("Unknown list type")
}

func (v *parameterVisitor) visitNode(ctx *ap.NodeContext) parameters.Node {
	return v.visitSingleNode(ctx.SingleNode().(*ap.SingleNodeContext))
}

func (v *parameterVisitor) visitParameter(ctx *ap.ParameterContext) parameters.Parameter {
	if ctx.SingleNode() != nil {
		return v.visitSingleNode(ctx.SingleNode().(*ap.SingleNodeContext))
	}
	if ctx.Node() != nil {
		return v.visitNode(ctx.Node().(*ap.NodeContext))
	}
	if ctx.Piece() != nil {
		return v.visitPiece(ctx.Piece().(*ap.PieceContext))
	}
	if ctx.Position() != nil {
		return v.visitPosition(ctx.Position().(*ap.PositionContext))
	}
	if ctx.Coord() != nil {
		return v.visitCoord(ctx.Coord().(*ap.CoordContext))
	}
	if ctx.List() != nil {
		return v.visitList(ctx.List().(*ap.ListContext))
	}
	if v.functionalEnabled && ctx.QUESTIONMARK() != nil {
		return parameters.PlaceHolder{}
	}
	panic(fmt.Sprintf("Unknown parameter type: %v", ctx.GetText()))
}

func (v *parameterVisitor) visitPiece(ctx *ap.PieceContext) parameters.Piece {
	return *parameters.NewPiece(v.visitNode(ctx.Node().(*ap.NodeContext)))
}

func (v *parameterVisitor) visitPosition(ctx *ap.PositionContext) parameters.Position {
	return *parameters.NewPosition(v.visitNode(ctx.Node().(*ap.NodeContext)))
}

func (v *parameterVisitor) visitSides(ctx *ap.SidesContext) []string {
	sides := make([]string, len(ctx.AllSide()))
	for i, side := range ctx.AllSide() {
		sides[i] = side.GetText()
	}
	return sides
}

func (v *parameterVisitor) visitSingleNode(ctx *ap.SingleNodeContext) parameters.Node {
	sides := v.visitSides(ctx.Sides().(*ap.SidesContext))
	typedSides := make([]parameters.Side, len(sides))
	for i, side := range sides {
		typedSides[i] = v.sideConstructor(side)
	}
	return *parameters.NewNode(typedSides)
}
