package generatorvisitor

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
)

type Visitor struct {
	size int
}

func NewVisitor() Visitor {
	return Visitor{}
}

func (v *Visitor) Visit(tree antlr.Tree) models.Puzzle {
	ctx, ok := tree.(*cp.ConfigFileContext)
	if !ok {
		panic("Invalid tree")
	}
	return v.visitConfigFile(ctx)
}

func (v *Visitor) visitConfigFile(ctx *cp.ConfigFileContext) models.Puzzle {
	for _, line := range ctx.AllConfigLine() {
		stateDef := line.(*cp.ConfigLineContext).StateDef()
		if stateDef != nil {
			return v.visitStateDef(stateDef.(*cp.StateDefContext))
		}
		sizeDef := line.(*cp.ConfigLineContext).SizeDef()
		if sizeDef != nil {
			v.visitSizeDef(sizeDef.(*cp.SizeDefContext))
		}
	}
	panic("No state definition found!")
}

func (v *Visitor) visitSizeDef(ctx *cp.SizeDefContext) {
	size := ctx.NUMBER().GetText()
	v.size, _ = strconv.Atoi(size)
}

func (v *Visitor) visitState(ctx *cp.StateContext) models.Puzzle {
	if ctx.RANDOM() != nil {
		panic("random not implemented yet!")
	}
	if ctx.AdvancedState() != nil {
		visitor := newAdvancedStateVisitor()
		return visitor.visitAdvancedState(ctx.AdvancedState().(*cp.AdvancedStateContext))
	}
	visitor := newBeginnerStateVisitor(v.size)
	return visitor.visitBeginnerState(ctx.BeginnerState().(*cp.BeginnerStateContext))
}

func (v *Visitor) visitStateDef(ctx *cp.StateDefContext) models.Puzzle {
	return v.visitState(ctx.State().(*cp.StateContext))
}
