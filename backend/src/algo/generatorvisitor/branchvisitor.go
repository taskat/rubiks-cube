package generatorvisitor

import (
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
)

func (v *Visitor) visitBranch(ctx *ap.BranchContext) {
	if ctx.IfBranch() != nil {
		v.visitIfBranch(ctx.IfBranch().(*ap.IfBranchContext))
		return
	}
	v.visitPrepareBranch(ctx.PrepareBranch().(*ap.PrepareBranchContext))
}

func (v *Visitor) visitBranches(ctx *ap.BranchesContext) {
	for _, branch := range ctx.AllBranch() {
		v.visitBranch(branch.(*ap.BranchContext))
	}
}

func (v *Visitor) visitConsecutive(ctx *ap.ConsecutiveContext) int {
	consec, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return consec
}

func (v *Visitor) visitIfBranch(ctx *ap.IfBranchContext) {
	conditionFunc := v.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))
	condition := algorithm.NewCondition(conditionFunc)
	v.nextSetter(condition)
	trueAction := v.visitDoDef(ctx.DoDef().(*ap.DoDefContext), condition.TrueSetter())
	trueAction.NextSetter()(v.currentGoal)
	v.nextSetter = condition.FalseSetter()
}

func (v *Visitor) visitPrepareBranch(ctx *ap.PrepareBranchContext) {
	if ctx.Algorithm() != nil {
		action := v.visitAlgorithmBlock(ctx.Algorithm().(*ap.AlgorithmContext), v.nextSetter)
		action.NextSetter()(v.currentGoal)
		v.nextSetter = v.currentGoal.TrueSetter()
		return
	}
	action := v.visitDoDef(ctx.DoDef().(*ap.DoDefContext), func(block algorithm.Block) {})
	consec := v.visitConsecutive(ctx.Consecutive().(*ap.ConsecutiveContext))
	prep := algorithm.NewPrepare(action, consec)
	v.nextSetter(prep)
	v.nextSetter = prep.NextSetter()
}
