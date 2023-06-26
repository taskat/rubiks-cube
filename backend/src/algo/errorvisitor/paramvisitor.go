package errorvisitor

import (
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/basevisitor"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type paramVisitor struct {
	basevisitor.ErrorVisitor
	sides       []string
	elementType iType
}

func newParamVisitor(errorVisitor basevisitor.ErrorVisitor, sides []string, elementType iType) *paramVisitor {
	return &paramVisitor{
		ErrorVisitor: errorVisitor,
		sides:        sides,
		elementType:  elementType,
	}
}

func (v *paramVisitor) visitCoord(ctx *ap.CoordContext) iType {
	v.visitSide(ctx.Side().(*ap.SideContext))
	rowString := ctx.NUMBER(0).GetText()
	row, err := strconv.Atoi(rowString)
	if err != nil {
		panic("Invalid row")
	}
	colString := ctx.NUMBER(1).GetText()
	col, err := strconv.Atoi(colString)
	if err != nil {
		panic("Invalid col")
	}
	if row < 1 || row > 3 {
		errCtx := eh.NewContext(ctx.NUMBER(0).GetSymbol().GetLine(), ctx.NUMBER(0).GetSymbol().GetColumn())
		v.Eh().AddError(errCtx, "Row must be between 1 and 3", v.FileName())
	}
	if col < 1 || col > 3 {
		errCtx := eh.NewContext(ctx.NUMBER(1).GetSymbol().GetLine(), ctx.NUMBER(1).GetSymbol().GetColumn())
		v.Eh().AddError(errCtx, "Col must be between 1 and 3", v.FileName())
	}
	return coord
}

func (v *paramVisitor) visitList(ctx *ap.ListContext) iType {
	if ctx.AllNode() != nil {
		for _, node := range ctx.AllNode() {
			v.visitNode(node.(*ap.NodeContext))
		}
		return newListType(node)
	}
	if ctx.AllPiece() != nil {
		for _, piece := range ctx.AllPiece() {
			v.visitPiece(piece.(*ap.PieceContext))
		}
		return newListType(piece)
	}
	if ctx.AllPosition() != nil {
		for _, position := range ctx.AllPosition() {
			v.visitPosition(position.(*ap.PositionContext))
		}
		return newListType(position)
	}
	if ctx.AllCoord() != nil {
		for _, coord := range ctx.AllCoord() {
			v.visitCoord(coord.(*ap.CoordContext))
		}
		return newListType(coord)
	}
	panic("Invalid list")
}

func (v *paramVisitor) visitNode(ctx *ap.NodeContext) iType {
	return v.visitSingleNode(ctx.SingleNode().(*ap.SingleNodeContext))
}

func (v *paramVisitor) visitParameter(ctx *ap.ParameterContext) iType {
	switch {
	case ctx.SingleNode() != nil:
		return v.visitSingleNode(ctx.SingleNode().(*ap.SingleNodeContext))
	case ctx.Node() != nil:
		return v.visitNode(ctx.Node().(*ap.NodeContext))
	case ctx.Piece() != nil:
		return v.visitPiece(ctx.Piece().(*ap.PieceContext))
	case ctx.Position() != nil:
		return v.visitPosition(ctx.Position().(*ap.PositionContext))
	case ctx.Coord() != nil:
		return v.visitCoord(ctx.Coord().(*ap.CoordContext))
	case ctx.List() != nil:
		return v.visitList(ctx.List().(*ap.ListContext))
	case ctx.QUESTIONMARK() != nil:
		if v.elementType == nil {
			v.Eh().AddError(ctx, "'?' can only appear in functional expression", v.FileName())
			//todo return type
		}
	}
	panic("Invalid parameter")
}

func (v *paramVisitor) visitPiece(ctx *ap.PieceContext) iType {
	v.visitNode(ctx.Node().(*ap.NodeContext))
	return piece
}

func (v *paramVisitor) visitPosition(ctx *ap.PositionContext) iType {
	v.visitNode(ctx.Node().(*ap.NodeContext))
	return position
}

func (v *paramVisitor) visitSide(ctx *ap.SideContext) {
	for _, side := range v.sides {
		if side == ctx.WORD().GetText() {
			return
		}
	}
	v.Eh().AddError(ctx, "Unknown side: "+ctx.WORD().GetText(), v.FileName())
}

func (v *paramVisitor) visitSides(ctx *ap.SidesContext) {
	for _, side := range ctx.AllSide() {
		v.visitSide(side.(*ap.SideContext))
	}
}

func (v *paramVisitor) visitSingleNode(ctx *ap.SingleNodeContext) iType {
	v.visitSides(ctx.Sides().(*ap.SidesContext))
	if ctx.NUMBER() != nil {
		v.Eh().AddError(ctx, "Unexpected number", v.FileName())
	}
	return node
}
