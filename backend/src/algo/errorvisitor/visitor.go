package errorvisitor

import (
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

func (v *Visitor) visitAlgorithmFile(ctx *ap.AlgorithmFileContext) {
	v.table.PushScope(scope.NewErrorScope())
	if ctx.Helpers() != nil {
		v.visitHelpers(ctx.Helpers().(*ap.HelpersContext))
	}
	v.table.CheckForErrorsTop(v.eh, v.fileName)
	v.visitSteps(ctx.Steps().(*ap.StepsContext))
	v.table.PopScope()
}

func (v *Visitor) visitBranches(ctx *ap.BranchesContext) {}

func (v *Visitor) visitDoDef(ctx *ap.DoDefContext) {}

func (v *Visitor) visitHelpers(ctx *ap.HelpersContext) {
	for _, helper := range ctx.AllHelperLine() {
		v.visitHelperLine(helper.(*ap.HelperLineContext))
	}
}

func (v *Visitor) visitGoal(ctx *ap.GoalContext) {
}

func (v *Visitor) visitHelperLine(ctx *ap.HelperLineContext) {
	visitor := newAlgorithmVisitor()
	visitor.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
	helperName := ctx.WORD().GetText()
	v.table.AddIdentifier(helperName, ctx)
}

func (v *Visitor) visitRuns(ctx *ap.RunsContext) {}

func (v *Visitor) visitStep(ctx *ap.StepContext) {
	v.table.PushScope(scope.NewErrorScope())
	goalDefs := getLines[*ap.GoalContext](ctx.AllStepLine())
	goalDef := checkOnlyOneDef(goalDefs, "goal definition", true, v)
	if goalDef != nil {
		v.visitGoal(*goalDef)
	}
	helpersDefs := getLines[*ap.HelperLineContext](ctx.AllStepLine())
	helperDef := checkOnlyOneDef(helpersDefs, "helper definition", false, v)
	if helperDef != nil {
		v.visitHelperLine(*helperDef)
	}
	doDefs := getLines[*ap.DoDefContext](ctx.AllStepLine())
	doDef := checkOnlyOneDef(doDefs, "do definition", false, v)
	if doDef != nil {
		v.visitDoDef(*doDef)
		runsDefs := getLines[*ap.RunsContext](ctx.AllStepLine())
		checkZeroDef(runsDefs, "runs definition", v)
		branchesDefs := getLines[*ap.BranchesContext](ctx.AllStepLine())
		checkZeroDef(branchesDefs, "branches definition", v)
	} else {
		runsDefs := getLines[*ap.RunsContext](ctx.AllStepLine())
		runsDef := checkOnlyOneDef(runsDefs, "runs definition", true, v)
		if runsDef != nil {
			v.visitRuns(*runsDef)
		}
		branchesDefs := getLines[*ap.BranchesContext](ctx.AllStepLine())
		branchesDef := checkOnlyOneDef(branchesDefs, "branches definition", true, v)
		if branchesDef != nil {
			v.visitBranches(*branchesDef)
		}
	}
	v.table.CheckForErrorsTop(v.eh, v.fileName)
	v.table.PopScope()
}

func (v *Visitor) visitSteps(ctx *ap.StepsContext) {
	for _, step := range ctx.AllStep() {
		v.visitStep(step.(*ap.StepContext))
	}
	if len(ctx.AllStep()) == 0 {
		v.eh.AddError(ctx, "No steps defined", v.fileName)
	}
}

func checkZeroDef[def antlr.ParserRuleContext](defs []def, defType string, v *Visitor) *def {
	if len(defs) != 0 {
		for _, d := range defs {
			v.eh.AddWarning(d, defType+" will be ignored", v.fileName)
		}
	}
	return nil
}

func checkOnlyOneDef[def antlr.ParserRuleContext](defs []def, defType string, necessary bool, v *Visitor) *def {
	if len(defs) > 1 {
		for _, d := range defs {
			v.eh.AddError(d, "Multiple "+defType+" found", v.fileName)
		}
		return nil
	} else if necessary && len(defs) == 0 {
		ctx := eh.NewContext(-1, -1)
		v.eh.AddError(ctx, "No "+defType+" found", v.fileName)
		return nil
	}
	if necessary || len(defs) == 1 {
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
