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
	if v.nextSetter != nil {
		v.nextSetter(algorithm.NewFinish())
	}
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
	visitor := newBoolExprVisitor(v.puzzle)
	return visitor.visitBoolExpr(ctx)
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
	//helpers
	for _, stepLine := range ctx.AllStepLine() {
		if stepLine.Helpers() != nil {
			v.visitHelpers(stepLine.Helpers().(*ap.HelpersContext))
		}
	}
	//runs
	runs := -1
	for _, stepLine := range ctx.AllStepLine() {
		if stepLine.Runs() != nil {
			runs = v.visitRuns(stepLine.Runs().(*ap.RunsContext))
			break
		}
	}
	//goal
	goal := false
	for _, stepLine := range ctx.AllStepLine() {
		if stepLine.Goal() != nil {
			v.visitGoal(stepLine.Goal().(*ap.GoalContext), runs)
			goal = true
			if v.startBlock == nil {
				v.startBlock = v.currentGoal
			} else {
				v.nextSetter(v.currentGoal)
			}
			v.nextSetter = v.currentGoal.FalseSetter()
			defer func() {
				v.nextSetter = v.currentGoal.TrueSetter()
			}()
			if v.previousBlock == nil {
				v.previousBlock = v.currentGoal
			}
			break
		}
	}
	//dodef
	for _, stepLine := range ctx.AllStepLine() {
		if stepLine.DoDef() != nil && !goal {
			var setter algorithm.Setter
			if v.nextSetter != nil {
				setter = v.nextSetter
			} else {
				setter = func(block algorithm.Block) {}
			}
			action := v.visitDoDef(stepLine.DoDef().(*ap.DoDefContext), setter)
			if v.currentGoal == nil {
				v.startBlock = action
			}
			v.previousBlock = action
			v.nextSetter = action.NextSetter()
			return
		}
		if stepLine.DoDef() != nil {
			action := v.visitDoDef(stepLine.DoDef().(*ap.DoDefContext), v.currentGoal.FalseSetter())
			action.NextSetter()(v.currentGoal)
			return
		}
	}
	//branches
	for _, stepLine := range ctx.AllStepLine() {
		if stepLine.Branches() != nil {
			v.visitBranches(stepLine.Branches().(*ap.BranchesContext))
		}
	}
}

func (v *Visitor) visitSteps(ctx *ap.StepsContext) {
	for _, step := range ctx.AllStep() {
		v.visitStep(step.(*ap.StepContext))
	}
}
