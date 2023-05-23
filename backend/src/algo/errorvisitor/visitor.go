package errorvisitor

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/symboltable"
	"github.com/taskat/rubiks-cube/src/symboltable/scope"
)

type Visitor struct {
	fileName string
	eh       *eh.Errorhandler
	table    *symboltable.Table[eh.IContext]
}

func NewVisitor(fileName string, errorHandler *eh.Errorhandler) *Visitor {
	table := symboltable.NewTable[eh.IContext]()
	table.PushScope(scope.NewErrorScope())
	return &Visitor{fileName: fileName, eh: errorHandler,
		table: table}
}

func (v *Visitor) Visit(tree antlr.Tree) {
	ctx, ok := tree.(*ap.AlgorithmFileContext)
	if !ok {
		panic("Invalid tree")
	}
	v.visitAlgorithmFile(ctx)
}

func (v *Visitor) visitAlgorithm(ctx *ap.AlgorithmContext) {
	visitor := newAlgorithmVisitor()
	visitor.visitAlgorithm(ctx)
}

func (v *Visitor) visitAlgorithmFile(ctx *ap.AlgorithmFileContext) {
	v.table.PushScope(scope.NewErrorScope())
	defer v.table.PopScope()
	if ctx.Helpers() != nil {
		v.visitHelpers(ctx.Helpers().(*ap.HelpersContext))
	}
	v.table.CheckForErrorsTop(v.eh, v.fileName)
	v.visitSteps(ctx.Steps().(*ap.StepsContext))
}

func (v *Visitor) visitBranches(ctx *ap.BranchesContext) {
	for _, branch := range ctx.AllBranch() {
		visitor := newBranchVisitor()
		visitor.visitBranch(branch.(*ap.BranchContext))
	}
}

func (v *Visitor) visitDoDef(ctx *ap.DoDefContext) {
	v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
}

func (v *Visitor) visitGoal(ctx *ap.GoalContext) {
	visitor := newBoolExprVisitor()
	visitor.visitBoolExpr(ctx.BoolExpr().(*ap.BoolExprContext))
}

func (v *Visitor) visitHelpers(ctx *ap.HelpersContext) {
	for _, helper := range ctx.AllHelperLine() {
		v.visitHelperLine(helper.(*ap.HelperLineContext))
	}
}

func (v *Visitor) visitHelperLine(ctx *ap.HelperLineContext) {
	v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
	helperName := ctx.WORD().GetText()
	v.table.AddIdentifier(helperName, ctx)
}

func (v *Visitor) visitRuns(ctx *ap.RunsContext) {
	runsString := ctx.NUMBER().GetText()
	_, err := strconv.Atoi(runsString)
	if err != nil {
		v.eh.AddError(eh.NewContext(ctx.NUMBER().GetSymbol().GetLine(), ctx.NUMBER().GetSymbol().GetColumn()), "number of runs has to be an integer", v.fileName)
	}
}

func (v *Visitor) visitStep(ctx *ap.StepContext) {
	v.table.PushScope(scope.NewErrorScope())
	defer func() {
		v.table.CheckForErrorsTop(v.eh, v.fileName)
		v.table.PopScope()
	}()
	if len(ctx.AllStepLine()) == 1 {
		visitDef(ctx, "do defintion", v, checkOneDef[*ap.DoDefContext], v.visitDoDef)
		return
	}
	visitDef(ctx, "goal definition", v, checkOneDef[*ap.GoalContext], v.visitGoal)
	visitDef(ctx, "runs definition", v, checkOneDef[*ap.RunsContext], v.visitRuns)
	visitDef(ctx, "helpers definition", v, checkOptionalDef[*ap.HelpersContext], v.visitHelpers)
	defs := make([]antlr.ParserRuleContext, 0)
	doDefs := getLines[*ap.DoDefContext](ctx.AllStepLine())
	for _, doDef := range doDefs {
		defs = append(defs, doDef)
	}
	branchesDefs := getLines[*ap.BranchesContext](ctx.AllStepLine())
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
		v.eh.AddError(ctx, "No steps defined", v.fileName)
	}
}

func visitDef[defType antlr.ParserRuleContext](ctx *ap.StepContext, defName string, v *Visitor,
	checkDefs func([]defType, string, *Visitor, eh.IContext) *defType, visitGoal func(defType)) *defType {
	defs := getLines[defType](ctx.AllStepLine())
	def := checkDefs(defs, defName, v, ctx)
	if def != nil {
		visitGoal(*def)
		return def
	}
	return nil
}

func visitOrDef[defType1, defType2 antlr.ParserRuleContext](defs []antlr.ParserRuleContext, ctx *ap.StepContext, defName string, v *Visitor,
	visitGoal1 func(defType1), visitGoal2 func(defType2)) antlr.ParserRuleContext {

	def := checkOneDef(defs, defName, v, ctx)
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

func checkZeroDef[def antlr.ParserRuleContext](defs []def, defType string, v *Visitor, parentCtx eh.IContext) *def {
	if len(defs) != 0 {
		for _, d := range defs {
			v.eh.AddWarning(d, defType+" will be ignored", v.fileName)
		}
	}
	return nil
}

func checkOneDef[def antlr.ParserRuleContext](defs []def, defType string, v *Visitor, parentCtx eh.IContext) *def {
	switch {
	case len(defs) > 1:
		for _, d := range defs {
			v.eh.AddError(d, "Multiple "+defType+" found", v.fileName)
		}
		return nil
	case len(defs) == 0:
		v.eh.AddError(parentCtx, "No "+defType+" found", v.fileName)
		return nil
	}
	return &defs[0]
}

func checkOptionalDef[def antlr.ParserRuleContext](defs []def, defType string, v *Visitor, parentCtx eh.IContext) *def {
	fmt.Println(defs)
	if len(defs) > 1 {
		for _, d := range defs {
			v.eh.AddError(d, "Multiple "+defType+" found", v.fileName)
		}
		return nil
	}
	if len(defs) == 1 {
		return &defs[0]
	}
	return nil
}

func getLines[def any](lines []ap.IStepLineContext) []def {
	result := make([]def, 0, 1)
	for _, l := range lines {
		line := l.(*ap.StepLineContext).GetChild(0)
		if line, ok := line.(def); ok {
			result = append(result, line)
		}
	}
	return result
}
