package generatorvisitor

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/symboltable"
	"github.com/taskat/rubiks-cube/src/symboltable/scope"
)

type Visitor struct {
	puzzle        models.Puzzle
	startBlock    algorithm.Block
	turns         *symboltable.Table[[]string]
	currentGoal   *algorithm.Goal
	previousBlock algorithm.Block
	nextSetter    algorithm.Setter
}

func NewVisitor(puzzle models.Puzzle) *Visitor {
	v := &Visitor{puzzle: puzzle, startBlock: nil, currentGoal: nil,
		previousBlock: nil, nextSetter: nil}
	v.initTurns()
	return v
}

func (v *Visitor) initTurns() {
	v.turns = symboltable.NewTable[[]string]()
	v.turns.PushScope(scope.NewScope[[]string]())
	for _, turn := range v.puzzle.GetConstraint().Turns {
		v.turns.AddIdentifier(turn, []string{turn})
	}
}

func (v *Visitor) Visit(tree antlr.Tree) algorithm.Algorithm {
	ctx, ok := tree.(*ap.AlgorithmFileContext)
	if !ok {
		panic("Invalid tree")
	}
	return v.visitAlgoFile(ctx)
}

func (v *Visitor) visitAlgoFile(ctx *ap.AlgorithmFileContext) algorithm.Algorithm {
	v.turns.PushScope(scope.NewScope[[]string]())
	defer v.turns.PopScope()
	if ctx.Helpers() != nil {
		v.visitHelpers(ctx.Helpers().(*ap.HelpersContext))
	}
	v.visitSteps(ctx.Steps().(*ap.StepsContext))
	return algorithm.NewAlgorithm(v.startBlock)
}

func (v *Visitor) visitAlgorithm(ctx *ap.AlgorithmContext) []string {
	visitor := newAlgorithmVisitor(v.turns)
	return visitor.visitAlgorithm(ctx)
}

func (v *Visitor) visitAlgorithmBlock(ctx *ap.AlgorithmContext, setter algorithm.Setter) *algorithm.Action {
	visitor := newAlgorithmBlockVisitor(v.turns)
	action := visitor.visitAlgorithm(ctx)
	setter(action)
	return action
}

func (v *Visitor) visitBoolExpr(ctx *ap.BoolExprContext) algorithm.ConditionFunc {
	return nil
}

func (v *Visitor) visitDoDef(ctx *ap.DoDefContext, setter algorithm.Setter) *algorithm.Action {
	return v.visitAlgorithmBlock(ctx.Algorithm().(*ap.AlgorithmContext), setter)
}

func (v *Visitor) visitGoal(ctx *ap.GoalContext, runs int) {
	conditionFunc := v.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))
	v.currentGoal = algorithm.NewGoal(conditionFunc, runs)
}

func (v *Visitor) visitHelperLine(ctx *ap.HelperLineContext) {
	moves := v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
	name := ctx.WORD().GetText()
	v.turns.AddIdentifier(name, moves)
}

func (v *Visitor) visitHelpers(ctx *ap.HelpersContext) {
	for _, helper := range ctx.AllHelperLine() {
		v.visitHelperLine(helper.(*ap.HelperLineContext))
	}
}

func (v *Visitor) visitRuns(ctx *ap.RunsContext) int {
	runs, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return runs
}

func (v *Visitor) visitStep(ctx *ap.StepContext) {
	v.turns.PushScope(scope.NewScope[[]string]())
	defer v.turns.PopScope()
	for _, stepLine := range ctx.AllStepLine() {
		v.visitStepLine(stepLine.(*ap.StepLineContext))
	}
}

func (v *Visitor) visitStepLine(ctx *ap.StepLineContext) {
	if ctx.DoDef() != nil && ctx.Goal() == nil {
		var setter algorithm.Setter
		if v.currentGoal != nil {
			setter = v.currentGoal.FalseSetter()
		} else {
			setter = v.nextSetter
		}
		action := v.visitDoDef(ctx.DoDef().(*ap.DoDefContext), setter)
		if v.currentGoal == nil {
			v.startBlock = action
		}
		v.previousBlock = action
		v.nextSetter = action.NextSetter()
		return
	}
	runs := v.visitRuns(ctx.Runs().(*ap.RunsContext))
	v.visitGoal(ctx.Goal().(*ap.GoalContext), runs)
	if v.startBlock == nil {
		v.startBlock = v.currentGoal
	}
	if v.previousBlock != nil {
		v.nextSetter(v.currentGoal)
	} else {
		v.previousBlock = v.currentGoal
	}
	if ctx.DoDef() != nil {
		action := v.visitDoDef(ctx.DoDef().(*ap.DoDefContext), v.currentGoal.FalseSetter())
		action.NextSetter()(v.currentGoal)
		return
	}
	v.visitBranches(ctx.Branches().(*ap.BranchesContext))
}

func (v *Visitor) visitSteps(ctx *ap.StepsContext) {
	for _, step := range ctx.AllStep() {
		v.visitStep(step.(*ap.StepContext))
	}
}
