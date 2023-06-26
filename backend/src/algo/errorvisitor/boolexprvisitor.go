package errorvisitor

import (
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/basevisitor"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/symboltable"
)

type boolExprVisitor struct {
	basevisitor.ErrorVisitor
	operators       *symboltable.Table[operator]
	sides           []string
	nested          bool
	listElementType iType
}

func newBoolExprVisitor(errorVisitor basevisitor.ErrorVisitor, operators *symboltable.Table[operator],
	sides []string, nested bool) *boolExprVisitor {
	return &boolExprVisitor{ErrorVisitor: errorVisitor, operators: operators,
		sides: sides, nested: nested}
}

func (v *boolExprVisitor) visitBinaryExpr(ctx *ap.BinaryExprContext) {
	name := ctx.WORD().GetText()
	errCtx := eh.NewContext(ctx.WORD().GetSymbol().GetLine(), ctx.WORD().GetSymbol().GetColumn())
	operator, err := v.operators.GetIdentifier(name)
	if err != nil {
		v.Eh().AddError(errCtx, "Unknown operator: "+name, v.FileName())
		return
	}
	binary := operator.toBinary()
	if binary == nil {
		v.Eh().AddError(errCtx, "Non-binary operator "+name+" used as binary", v.FileName())
		return
	}
	leftParam := v.visitParameter(ctx.Parameter(0).(*ap.ParameterContext))
	if leftParam != nil {
		err := binary.acceptLeftParam(leftParam)
		if err != nil {
			v.Eh().AddError(ctx, err.Error(), v.FileName())
		}
	}
	rightParam := v.visitParameter(ctx.Parameter(1).(*ap.ParameterContext))
	if rightParam != nil {
		err := binary.acceptRightParam(rightParam)
		if err != nil {
			v.Eh().AddError(ctx, err.Error(), v.FileName())
		}
	}
}

func (v *boolExprVisitor) visitBoolExpr(ctx *ap.BoolExprContext) {
	switch {
	case ctx.UnaryOp() != nil:
		v.visitBoolExpr(ctx.BoolExpr(0).(*ap.BoolExprContext))
	case ctx.BinaryOp() != nil:
		v.visitBoolExpr(ctx.BoolExpr(0).(*ap.BoolExprContext))
		v.visitBoolExpr(ctx.BoolExpr(1).(*ap.BoolExprContext))
	case ctx.LPAREN() != nil:
		v.visitBoolExpr(ctx.BoolExpr(0).(*ap.BoolExprContext))
	case ctx.Expr() != nil:
		v.visitExpr(ctx.Expr().(*ap.ExprContext))
	}
}

func (v *boolExprVisitor) visitExpr(ctx *ap.ExprContext) {
	switch {
	case ctx.UnaryExpr() != nil:
		v.visitUnaryExpr(ctx.UnaryExpr().(*ap.UnaryExprContext))
	case ctx.BinaryExpr() != nil:
		v.visitBinaryExpr(ctx.BinaryExpr().(*ap.BinaryExprContext))
	case ctx.FunctionalExpr() != nil:
		if v.nested {
			v.Eh().AddError(ctx, "Nested functional expressions are not allowed", v.FileName())
			return
		}
		v.visitFunctionalExpr(ctx.FunctionalExpr().(*ap.FunctionalExprContext))
	}
}

func (v *boolExprVisitor) visitFunctionalExpr(ctx *ap.FunctionalExprContext) {
	paramVisitor := newParamVisitor(v.ErrorVisitor, v.sides, nil)
	v.listElementType = paramVisitor.visitList(ctx.List().(*ap.ListContext))
	visitor := newBoolExprVisitor(v.ErrorVisitor, v.operators, v.sides, true)
	visitor.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))

}

func (v *boolExprVisitor) visitParameter(ctx *ap.ParameterContext) iType {
	var visitor *paramVisitor
	if v.nested {
		visitor = newParamVisitor(v.ErrorVisitor, v.sides, v.listElementType)
	} else {
		visitor = newParamVisitor(v.ErrorVisitor, v.sides, nil)
	}
	return visitor.visitParameter(ctx)
}

func (v *boolExprVisitor) visitUnaryExpr(ctx *ap.UnaryExprContext) {
	name := ctx.WORD().GetText()
	errCtx := eh.NewContext(ctx.WORD().GetSymbol().GetLine(), ctx.WORD().GetSymbol().GetColumn())
	operator, err := v.operators.GetIdentifier(name)
	if err != nil {
		v.Eh().AddError(errCtx, "Unknown operator: "+name, v.FileName())
		return
	}
	unary := operator.toUnary()
	if unary == nil {
		v.Eh().AddError(errCtx, "Non-unary operator "+name+" used as unary", v.FileName())
		return
	}
	param := v.visitParameter(ctx.Parameter().(*ap.ParameterContext))
	if param != nil {
		err := unary.acceptParam(param)
		if err != nil {
			v.Eh().AddError(ctx, err.Error(), v.FileName())
		}
	}
}
