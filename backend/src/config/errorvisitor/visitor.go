package errorvisitor

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/taskat/rubiks-cube/src/basevisitor"
	"github.com/taskat/rubiks-cube/src/basevisitor/util"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Visitor struct {
	basevisitor.ErrorVisitor
	stateDescription    string
	size                int
	stateDescriptionCtx eh.IContext
}

func NewVisitor(fileName string, errorHandler *eh.Errorhandler) *Visitor {
	return &Visitor{ErrorVisitor: *basevisitor.NewErrorVisitor(errorHandler, fileName)}
}

func (v *Visitor) Visit(tree antlr.Tree) {
	ctx, ok := tree.(*cp.ConfigFileContext)
	if !ok {
		panic("Invalid tree")
	}
	v.visitConfigFile(ctx)
}

func (v *Visitor) visitConfigFile(ctx *cp.ConfigFileContext) {
	visitDef(ctx, "puzzle definition", v, util.CheckOneDef[*cp.PuzzleTypeDefContext], v.visitPuzzleTypeDef)
	visitDef(ctx, "size definition", v, util.CheckOneDef[*cp.SizeDefContext], v.visitSizeDef)
	visitDef(ctx, "state description definition", v, util.CheckOptionalDef[*cp.StateDescriptionDefContext], v.visitStateDescriptionDef)
	visitDef(ctx, "state definition", v, util.CheckOneDef[*cp.StateDefContext], v.visitStateDef)
}

func (v *Visitor) visitPuzzleTypeDef(ctx *cp.PuzzleTypeDefContext) {
	// Nothing to check
}

func (v *Visitor) visitSizeDef(ctx *cp.SizeDefContext) {
	sizeString := ctx.NUMBER().GetText()
	var err error
	v.size, err = strconv.Atoi(sizeString)
	min := 2
	max := 7
	if err != nil {
		v.Eh().AddError(ctx, "cannot convert to size (integer)", v.FileName())
	} else if v.size < min || v.size > max {
		v.Eh().AddError(ctx, fmt.Sprintf("Size has to be >=%d and <=%d", min, max), v.FileName())
	}
}

func (v *Visitor) visitState(ctx *cp.StateContext) {
	if ctx.RANDOM() != nil {
		if v.stateDescription != "" {
			v.Eh().AddInfo(v.stateDescriptionCtx, "state description can be omitted", v.FileName())
		}
	}
	errorMsg := "state is not consistent with state description type"
	if ctx.BeginnerState() != nil {
		if v.stateDescription != "beginner" {
			v.Eh().AddError(ctx.BeginnerState(), errorMsg, v.FileName())
			if v.stateDescriptionCtx != nil {
				v.Eh().AddError(v.stateDescriptionCtx, errorMsg, v.FileName())
			}
		} else {
			visitor := newBeginnerStateVisitor(v.FileName(), v.Eh(), v.size)
			visitor.visitBeginnerState(ctx.BeginnerState().(*cp.BeginnerStateContext))
		}
	}
	if ctx.AdvancedState() != nil {
		if v.stateDescription != "advanced" {
			v.Eh().AddError(ctx.AdvancedState(), errorMsg, v.FileName())
			if v.stateDescriptionCtx != nil {
				v.Eh().AddError(v.stateDescriptionCtx, errorMsg, v.FileName())
			}
		} else {
			visitor := newAdvancedStateVisitor(v.FileName(), v.Eh(), v.size)
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

func visitDef[defType antlr.ParserRuleContext](ctx *cp.ConfigFileContext, defName string, v util.Visitor,
	checkDefs func([]defType, string, util.Visitor, eh.IContext) *defType, visitGoal func(defType)) *defType {
	defs := util.GetLines[defType](ctx.AllConfigLine())
	def := checkDefs(defs, defName, v, ctx)
	if def != nil {
		visitGoal(*def)
		return def
	}
	return nil
}
