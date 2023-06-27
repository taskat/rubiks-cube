package errorvisitor

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/basevisitor"
	"github.com/taskat/rubiks-cube/src/basevisitor/util"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/symboltable"
	"github.com/taskat/rubiks-cube/src/symboltable/scope"
)

type Visitor struct {
	basevisitor.ErrorVisitor
	turns           *symboltable.Table[eh.IContext]
	constraint      *models.Constraint
	currentStepRuns int
	ts              *typeSystem
}

func NewVisitor(fileName string, errorHandler *eh.Errorhandler, constraint models.Constraint) *Visitor {
	v := Visitor{ErrorVisitor: *basevisitor.NewErrorVisitor(errorHandler, fileName),
		currentStepRuns: 0, constraint: &constraint, ts: newTypeSystem()}
	v.initTurns(constraint.Turns)
	return &v
}

func (v *Visitor) initTurns(turns []string) {
	v.turns = symboltable.NewTable[eh.IContext]()
	v.turns.PushScope(scope.NewErrorScope())
	for _, turn := range turns {
		v.turns.AddIdentifier(turn, nil)
	}
}

func (v *Visitor) Visit(tree antlr.Tree) {
	ctx, ok := tree.(*ap.AlgorithmFileContext)
	if !ok {
		panic("Invalid tree")
	}
	v.visitAlgorithmFile(ctx)
}

func (v *Visitor) visitAlgorithm(ctx *ap.AlgorithmContext) {
	visitor := newAlgorithmVisitor(v.ErrorVisitor, v.turns)
	visitor.visitAlgorithm(ctx)
}

func (v *Visitor) visitAlgorithmFile(ctx *ap.AlgorithmFileContext) {
	v.turns.PushScope(scope.NewErrorScope())
	defer v.turns.PopScope()
	if ctx.Helpers() != nil {
		v.visitHelpers(ctx.Helpers().(*ap.HelpersContext))
	}
	v.turns.CheckForErrorsTop(v.Eh(), v.FileName())
	v.visitSteps(ctx.Steps().(*ap.StepsContext))
}

func (v *Visitor) visitBoolExpr(ctx *ap.BoolExprContext) {
	visitor := newBoolExprVisitor(v, false)
	visitor.visitBoolExpr(ctx)
}

func (v *Visitor) visitBranches(ctx *ap.BranchesContext) {
	visitor := newBranchVisitor(v)
	visitor.visitBranches(ctx)
}

func (v *Visitor) visitDoDef(ctx *ap.DoDefContext) {
	v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
}

func (v *Visitor) visitGoal(ctx *ap.GoalContext) {
	v.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))
}

func (v *Visitor) visitHelpers(ctx *ap.HelpersContext) {
	for _, helper := range ctx.AllHelperLine() {
		v.visitHelperLine(helper.(*ap.HelperLineContext))
	}
}

func (v *Visitor) visitHelperLine(ctx *ap.HelperLineContext) {
	v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
	helperName := ctx.WORD().GetText()
	v.turns.AddIdentifier(helperName, ctx)
}

func (v *Visitor) visitRuns(ctx *ap.RunsContext) {
	runsString := ctx.NUMBER().GetText()
	runs, err := strconv.Atoi(runsString)
	if err != nil {
		errCtx := eh.NewContextFromTerminal(ctx.NUMBER())
		v.Eh().AddError(errCtx, "number of runs has to be an integer", v.FileName())
	}
	v.currentStepRuns = runs
}

func (v *Visitor) visitStep(ctx *ap.StepContext) {
	v.turns.PushScope(scope.NewErrorScope())
	defer func() {
		v.turns.CheckForErrorsTop(v.Eh(), v.FileName())
		v.turns.PopScope()
	}()
	if len(ctx.AllStepLine()) == 1 {
		visitDef(ctx, "do defintion", v, util.CheckOneDef[*ap.DoDefContext], v.visitDoDef)
		return
	}
	visitDef(ctx, "goal definition", v, util.CheckOneDef[*ap.GoalContext], v.visitGoal)
	visitDef(ctx, "runs definition", v, util.CheckOneDef[*ap.RunsContext], v.visitRuns)
	visitDef(ctx, "helpers definition", v, util.CheckOptionalDef[*ap.HelpersContext], v.visitHelpers)
	defs := make([]antlr.ParserRuleContext, 0)
	doDefs := util.GetLines[*ap.DoDefContext](ctx.AllStepLine())
	for _, doDef := range doDefs {
		defs = append(defs, doDef)
	}
	branchesDefs := util.GetLines[*ap.BranchesContext](ctx.AllStepLine())
	for _, branchesDef := range branchesDefs {
		defs = append(defs, branchesDef)
	}
	visitOrDef(defs, ctx, "do/branches definition", v, v.visitDoDef, v.visitBranches)
}

func (v *Visitor) visitSteps(ctx *ap.StepsContext) {
	for _, step := range ctx.AllStep() {
		v.visitStep(step.(*ap.StepContext))
	}
	if len(ctx.AllStep()) == 0 {
		v.Eh().AddError(ctx, "No steps defined", v.FileName())
	}
}

func visitDef[defType antlr.ParserRuleContext](ctx *ap.StepContext, defName string, v util.Visitor,
	checkDefs func([]defType, string, util.Visitor, eh.IContext) *defType, visitGoal func(defType)) *defType {
	defs := util.GetLines[defType](ctx.AllStepLine())
	def := checkDefs(defs, defName, v, ctx)
	if def != nil {
		visitGoal(*def)
		return def
	}
	return nil
}

func visitOrDef[defType1, defType2 antlr.ParserRuleContext](defs []antlr.ParserRuleContext, ctx *ap.StepContext, defName string, v util.Visitor,
	visitGoal1 func(defType1), visitGoal2 func(defType2)) antlr.ParserRuleContext {

	def := util.CheckOneDef(defs, defName, v, ctx)
	if def == nil {
		return nil
	}
	if type1, ok := (*def).(defType1); ok {
		visitGoal1(type1)
		return type1
	}
	if type2, ok := (*def).(defType2); ok {
		visitGoal2(type2)
		return type2
	}
	return nil
}
