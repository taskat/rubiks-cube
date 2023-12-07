package errorvisitor

import (
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/basevisitor"
)

type branchVisitor struct {
	basevisitor.ErrorVisitor
	parentVisitor *Visitor
	warningAll    bool
}

func newBranchVisitor(v *Visitor) *branchVisitor {
	return &branchVisitor{ErrorVisitor: v.ErrorVisitor, parentVisitor: v,
		warningAll: false}
}

func (v *branchVisitor) visitBranch(ctx *ap.BranchContext, last bool) {
	if ctx.IfBranch() != nil {
		v.visitIfBranch(ctx.IfBranch().(*ap.IfBranchContext), last)
	} else if ctx.PrepareBranch() != nil {
		v.visitPrepareBranch(ctx.PrepareBranch().(*ap.PrepareBranchContext), last)
	}
}

func (v *branchVisitor) visitBranches(ctx *ap.BranchesContext) {
	for i, branch := range ctx.AllBranch() {
		if v.warningAll {
			v.Eh().AddWarning(branch, "Branch will be ignored", v.FileName())
		}
		v.visitBranch(branch.(*ap.BranchContext), i == len(ctx.AllBranch())-1)
	}
}

func (v *branchVisitor) visitConsecutive(ctx *ap.ConsecutiveContext) {
	numberString := ctx.NUMBER().GetText()
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	if number >= v.parentVisitor.currentStepRuns {
		v.Eh().AddInfo(ctx, "Consecutive runs more than step runs", v.FileName())
	}
}

func (v *branchVisitor) visitIfBranch(ctx *ap.IfBranchContext, last bool) {
	v.parentVisitor.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))
	v.parentVisitor.visitDoDef(ctx.DoDef().(*ap.DoDefContext))
	if last {
		parentCtx := ctx.GetParent().GetParent().(*ap.BranchesContext)
		v.Eh().AddError(parentCtx, "Last branch is not a prepare branch", v.FileName())
	}
}

func (v *branchVisitor) visitPrepareBranch(ctx *ap.PrepareBranchContext, last bool) {
	if ctx.Algorithm() != nil {
		v.parentVisitor.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
		v.warningAll = true
	} else {
		v.parentVisitor.visitDoDef(ctx.DoDef().(*ap.DoDefContext))
		v.visitConsecutive(ctx.Consecutive().(*ap.ConsecutiveContext))
	}
}
