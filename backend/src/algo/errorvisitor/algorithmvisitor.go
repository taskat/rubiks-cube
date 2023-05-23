package errorvisitor

import (
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
)

type algorithmVisitor struct {
}

func newAlgorithmVisitor() *algorithmVisitor {
	return &algorithmVisitor{}
}

func (v *algorithmVisitor) visitAlgorithm(ctx *ap.AlgorithmContext) {

}
