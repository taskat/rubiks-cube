package generatorvisitor

import (
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type boolExprVisitor struct {
	constraint models.Constraint
}

func newBoolExprVisitor(c models.Constraint) *boolExprVisitor {
	return &boolExprVisitor{constraint: c}
}

func (v *boolExprVisitor) buildAtFunc(leftParam, rightParam parameters.Parameter) algorithm.ConditionFunc {
	var left parameters.Piece
	var right parameters.Position
	switch typedLeftParam := leftParam.(type) {
	case parameters.Piece:
		left = typedLeftParam
	case parameters.Node:
		left = *parameters.NewPiece(typedLeftParam)
	}
	switch typedRightParam := rightParam.(type) {
	case parameters.Position:
		right = typedRightParam
	case parameters.Node:
		right = *parameters.NewPosition(typedRightParam)
	}
	return at(left, right)
}

func (v *boolExprVisitor) buildBinaryFuncFromUnary(param parameters.Parameter, builder func(leftParam, rightParam parameters.Parameter) algorithm.ConditionFunc) algorithm.ConditionFunc {
	if list, ok := param.(parameters.List[parameters.Parameter]); ok {
		return func(p models.Puzzle) bool {
			for _, param := range list {
				if !builder(param, param)(p) {
					return false
				}
			}
			return true
		}
	}
	return builder(param, param)
}

func (v *boolExprVisitor) buildLikeFunc(leftParam, rightParam parameters.Parameter) algorithm.ConditionFunc {
	var left parameters.Piece
	var right parameters.Position
	switch typedLeftParam := leftParam.(type) {
	case parameters.Piece:
		left = typedLeftParam
	case parameters.Node:
		left = *parameters.NewPiece(typedLeftParam)
	}
	switch typedRightParam := rightParam.(type) {
	case parameters.Position:
		right = typedRightParam
	case parameters.Node:
		right = *parameters.NewPosition(typedRightParam)
	}
	return like(left, right)
}

func (v *boolExprVisitor) visitBinaryExpr(ctx *ap.BinaryExprContext) algorithm.ConditionFunc {
	left := v.visitParameter(ctx.Parameter(0).(*ap.ParameterContext))
	right := v.visitParameter(ctx.Parameter(1).(*ap.ParameterContext))
	switch ctx.WORD().GetText() {
	case "at":
		return v.buildAtFunc(left, right)
	case "like":
		return v.buildLikeFunc(left, right)
	}
	panic("Unknown binary expression")
}

func (v *boolExprVisitor) visitBinaryOp(ctx *ap.BinaryOpContext, left, right algorithm.ConditionFunc) algorithm.ConditionFunc {
	if ctx.AND() != nil {
		return func(p models.Puzzle) bool {
			return left(p) && right(p)
		}
	}
	if ctx.OR() != nil {
		return func(p models.Puzzle) bool {
			return left(p) || right(p)
		}
	}
	panic("Unknown binary operator")
}

func (v *boolExprVisitor) visitBoolExpr(ctx *ap.BoolExprContext) algorithm.ConditionFunc {
	if ctx.UnaryOp() != nil {
		condFunc := v.visitBoolExpr(ctx.BoolExpr(0).(*ap.BoolExprContext))
		return v.visitUnaryOp(ctx.UnaryOp().(*ap.UnaryOpContext), condFunc)
	}
	if ctx.BinaryOp() != nil {
		left := v.visitBoolExpr(ctx.BoolExpr(0).(*ap.BoolExprContext))
		right := v.visitBoolExpr(ctx.BoolExpr(1).(*ap.BoolExprContext))
		return v.visitBinaryOp(ctx.BinaryOp().(*ap.BinaryOpContext), left, right)
	}
	if ctx.LPAREN() != nil {
		return v.visitBoolExpr(ctx.BoolExpr(0).(*ap.BoolExprContext))
	}
	if ctx.Expr() != nil {
		return v.visitExpr(ctx.Expr().(*ap.ExprContext))
	}
	panic("Unknown boolean expression")
}

func (v *boolExprVisitor) visitExpr(ctx *ap.ExprContext) algorithm.ConditionFunc {
	if ctx.UnaryExpr() != nil {
		return v.visitUnaryExpr(ctx.UnaryExpr().(*ap.UnaryExprContext))
	}
	if ctx.BinaryExpr() != nil {
		return v.visitBinaryExpr(ctx.BinaryExpr().(*ap.BinaryExprContext))
	}
	if ctx.FunctionalExpr() != nil {
		return v.visitFunctionalExpr(ctx.FunctionalExpr().(*ap.FunctionalExprContext))
	}
	panic("Unknown expression")
}

func (v *boolExprVisitor) visitFunctionalExpr(ctx *ap.FunctionalExprContext) algorithm.ConditionFunc {
	return nil
}

func (v *boolExprVisitor) visitParameter(ctx *ap.ParameterContext) parameters.Parameter {
	visitor := newParameterVisitor()
	return visitor.visitParameter(ctx)
}

func (v *boolExprVisitor) visitUnaryExpr(ctx *ap.UnaryExprContext) algorithm.ConditionFunc {
	param := v.visitParameter(ctx.Parameter().(*ap.ParameterContext))
	switch ctx.WORD().GetText() {
	case "place":
		return v.buildBinaryFuncFromUnary(param, v.buildAtFunc)
	case "orientation":
		return v.buildBinaryFuncFromUnary(param, v.buildLikeFunc)
	}
	expectedColor := color.Color(ctx.WORD().GetText())
	switch typedParam := param.(type) {
	case parameters.Coord:
		return singleColorMatch(expectedColor, typedParam)
	case parameters.List[parameters.Coord]:
		return colorListMatch(expectedColor, typedParam)
	}
	panic("Unknown unary expression")
}

func (v *boolExprVisitor) visitUnaryOp(ctx *ap.UnaryOpContext, condFunc algorithm.ConditionFunc) algorithm.ConditionFunc {
	if ctx.NOT() != nil {
		return func(p models.Puzzle) bool {
			return !condFunc(p)
		}
	}
	panic("Unknown unary operator")
}
