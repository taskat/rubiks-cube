package errorvisitor

import (
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
)

type boolExprVisitor struct{}

func newBoolExprVisitor() *boolExprVisitor {
	return &boolExprVisitor{}
}

func (v *boolExprVisitor) visitBoolExpr(ctx *ap.BoolExprContext) {}
