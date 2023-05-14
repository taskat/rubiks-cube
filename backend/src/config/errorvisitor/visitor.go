package errorvisitor

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Visitor struct {
	fileName            string
	stateDescription    string
	stateDescriptionCtx eh.IContext
	eh                  eh.Errorhandler
}

func NewVisitor(fileName string, errorHandler eh.Errorhandler) *Visitor {
	return &Visitor{fileName: fileName, eh: errorHandler}
}

func (v *Visitor) Visit(tree antlr.Tree) {
	ctx, ok := tree.(*cp.ConfigFileContext)
	if !ok {
		panic("Invalid tree")
	}
	v.visitConfigFile(ctx)
}

func (v *Visitor) visitConfigFile(ctx *cp.ConfigFileContext) {
	puzzleDefs := getLines[*cp.PuzzleTypeDefContext](ctx.AllConfigLine())
	puzzleDef := checkOnlyOneDef(puzzleDefs, "puzzle definition", true, v)
	if puzzleDef != nil {
		v.visitPuzzleTypeDef(*puzzleDef)
	}
	sizeDefs := getLines[*cp.SizeDefContext](ctx.AllConfigLine())
	sizeDef := checkOnlyOneDef(sizeDefs, "size definition", true, v)
	if sizeDef != nil {
		v.visitSizeDef(*sizeDef)
	}
	stateDescriptionDefs := getLines[*cp.StateDescriptionDefContext](ctx.AllConfigLine())
	stateDescriptionDef := checkOnlyOneDef(stateDescriptionDefs, "state description definition", false, v)
	if stateDescriptionDef != nil {
		v.visitStateDescriptionDef(*stateDescriptionDef)
	}
	stateDefs := getLines[*cp.StateDefContext](ctx.AllConfigLine())
	stateDef := checkOnlyOneDef(stateDefs, "state definition", true, v)
	if stateDef != nil {
		v.visitStateDef(*stateDef)
	}
}

func (v *Visitor) visitPuzzleTypeDef(ctx *cp.PuzzleTypeDefContext) {
	// Nothing to check
}

func (v *Visitor) visitSizeDef(ctx *cp.SizeDefContext) {
	sizeString := ctx.NUMBER().GetText()
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		v.eh.AddError(ctx, "cannot convert to size (integer)", v.fileName)
	} else if size != 3 {
		v.eh.AddError(ctx, "Size can only be 3", v.fileName)
	}
}

func (v *Visitor) visitState(ctx *cp.StateContext) {
	if ctx.RANDOM() != nil {
		if v.stateDescription != "" {
			v.eh.AddInfo(v.stateDescriptionCtx, "state description can be omitted", v.fileName)
		}
	}
	errorMsg := "state is not consistent with state description type"
	if ctx.BeginnerState() != nil {
		if v.stateDescription != "beginner" {
			v.eh.AddError(ctx.BeginnerState(), errorMsg, v.fileName)
			if v.stateDescriptionCtx != nil {
				v.eh.AddError(v.stateDescriptionCtx, errorMsg, v.fileName)
			}
		} else {
			visitor := newBeginnerStateVisitor(v.fileName, &v.eh)
			visitor.visitBeginnerState(ctx.BeginnerState().(*cp.BeginnerStateContext))
		}
	}
	if ctx.AdvancedState() != nil {
		if v.stateDescription != "advanced" {
			v.eh.AddError(ctx.AdvancedState(), errorMsg, v.fileName)
			if v.stateDescriptionCtx != nil {
				v.eh.AddError(v.stateDescriptionCtx, errorMsg, v.fileName)
			}
		} else {
			visitor := newAdvancedStateVisitor(v.fileName, &v.eh)
			visitor.visitAdvancedState(ctx.AdvancedState().(*cp.AdvancedStateContext))
		}
	}
}

func (v *Visitor) visitStateDescriptionDef(ctx *cp.StateDescriptionDefContext) {
	v.stateDescription = ctx.StateDescription().GetText()
	v.stateDescriptionCtx = ctx.StateDescription()
}

func (v *Visitor) visitStateDef(ctx *cp.StateDefContext) {
	v.visitState(ctx.State().(*cp.StateContext))
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

func getLines[def any](lines []cp.IConfigLineContext) []def {
	result := make([]def, 0, 1)
	for _, l := range lines {
		line := l.(*cp.ConfigLineContext).GetChild(0)
		if line, ok := line.(def); ok {
			result = append(result, line)
		}
	}
	return result
}
