package errorvisitor

import (
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
)

type branchVisitor struct {}

func newBranchVisitor() *branchVisitor {
	return &branchVisitor{}
}

func (v *branchVisitor) visitBranch(ctx *ap.BranchContext) {}