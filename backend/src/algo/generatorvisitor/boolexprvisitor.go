package generatorvisitor

import (
	"fmt"

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

func (v *boolExprVisitor) buildBinaryFunc(left, right parameters.Parameter, op string) algorithm.ConditionFunc {
	return nil
}

func (v *boolExprVisitor) visitBinaryExpr(ctx *ap.BinaryExprContext) algorithm.ConditionFunc {
	return nil
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
		return v.buildBinaryFunc(param, param, "at")
	case "orientation":
		return v.buildBinaryFunc(param, param, "like")
	}
	expectedColor := color.Color(ctx.WORD().GetText())
	switch typedParam := param.(type) {
	case parameters.Coord:
		return func(p models.Puzzle) bool {
			fmt.Println("Expected color", expectedColor)
			fmt.Println("Actual color", p.GetColor(typedParam))
			return p.GetColor(typedParam)[0] == expectedColor[0]
		}
	case parameters.List[parameters.Coord]:
		return func(p models.Puzzle) bool {
			for _, coord := range typedParam {
				if p.GetColor(coord) != expectedColor {
					return false
				}
			}
			return true
		}
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
