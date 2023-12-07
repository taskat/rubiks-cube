package generatorvisitor

import (
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/color"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type conditionBuilderFunc func(param parameters.Parameter) algorithm.ConditionFunc

type boolExprVisitor struct {
	constraint        models.Constraint
	sideConstructor   func(s string) parameters.Side
	functionalEnabled bool
	conditionBuilder  conditionBuilderFunc
}

func newBoolExprVisitor(p models.Puzzle) *boolExprVisitor {
	return &boolExprVisitor{constraint: p.GetConstraint(), sideConstructor: p.SideConstructor(), functionalEnabled: false}
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
	var conditionBuilder func(left, right parameters.Parameter) algorithm.ConditionFunc
	switch ctx.WORD().GetText() {
	case "at":
		conditionBuilder = v.buildAtFunc
	case "like":
		conditionBuilder = v.buildLikeFunc
	default:
		panic("Unknown binary expression")
	}
	switch {
	case left == parameters.PlaceHolder{} && right == parameters.PlaceHolder{}:
		v.conditionBuilder = func(param parameters.Parameter) algorithm.ConditionFunc {
			return conditionBuilder(param, param)
		}
	case left == parameters.PlaceHolder{}:
		v.conditionBuilder = func(param parameters.Parameter) algorithm.ConditionFunc {
			return conditionBuilder(param, right)
		}
	case right == parameters.PlaceHolder{}:
		v.conditionBuilder = func(param parameters.Parameter) algorithm.ConditionFunc {
			return conditionBuilder(left, param)
		}
	default:
		return conditionBuilder(left, right)
	}
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

func (v *boolExprVisitor) visitFunction(ctx *ap.FunctionContext) func(builder conditionBuilderFunc, list parameters.List[parameters.Parameter]) algorithm.ConditionFunc {
	switch {
	case ctx.ALL() != nil:
		return all
	case ctx.ANY() != nil:
		return any
	case ctx.NONE() != nil:
		return none
	default:
		panic("Unknown function")
	}
}

func (v *boolExprVisitor) visitFunctionalExpr(ctx *ap.FunctionalExprContext) algorithm.ConditionFunc {
	paramVisitor := newParameterVisitor(false, v.sideConstructor)
	list := paramVisitor.visitList(ctx.List().(*ap.ListContext))
	v.functionalEnabled = true
	v.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))
	v.functionalEnabled = false
	builder := v.visitFunction(ctx.Function().(*ap.FunctionContext))
	return builder(v.conditionBuilder, list)
}

func (v *boolExprVisitor) visitParameter(ctx *ap.ParameterContext) parameters.Parameter {
	visitor := newParameterVisitor(v.functionalEnabled, v.sideConstructor)
	return visitor.visitParameter(ctx)
}

func (v *boolExprVisitor) visitUnaryExpr(ctx *ap.UnaryExprContext) algorithm.ConditionFunc {
	param := v.visitParameter(ctx.Parameter().(*ap.ParameterContext))
	switch ctx.WORD().GetText() {
	case "place":
		return v.buildBinaryFuncFromUnary(param, v.buildAtFunc)
	case "orientation":
		return v.buildBinaryFuncFromUnary(param, v.buildLikeFunc)
	case "same":
		paramList := param.(parameters.List[parameters.Parameter])
		correctList := make(parameters.List[parameters.Coord], len(paramList))
		for i, param := range paramList {
			correctList[i] = param.(parameters.Coord)
		}
		return same(correctList)
	}
	expectedColor := color.Color(ctx.WORD().GetText())
	switch typedParam := param.(type) {
	case parameters.Coord:
		return singleColorMatch(expectedColor, typedParam)
	case parameters.List[parameters.Parameter]:
		var list parameters.List[parameters.Coord]
		for _, param := range typedParam {
			list = append(list, param.(parameters.Coord))
		}
		return colorListMatch(expectedColor, list)
	case parameters.PlaceHolder:
		v.conditionBuilder = func(param parameters.Parameter) algorithm.ConditionFunc {
			return singleColorMatch(expectedColor, param.(parameters.Coord))
		}
		return nil
	}
	panic("Unknown unary expression: " + ctx.GetText())
}

func (v *boolExprVisitor) visitUnaryOp(ctx *ap.UnaryOpContext, condFunc algorithm.ConditionFunc) algorithm.ConditionFunc {
	if ctx.NOT() != nil {
		return func(p models.Puzzle) bool {
			return !condFunc(p)
		}
	}
	panic("Unknown unary operator")
}
