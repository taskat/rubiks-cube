package generatorvisitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
)

type Visitor struct{}

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
		def := line.(*cp.ConfigLineContext).StateDef()
		if def != nil {
			return v.visitStateDef(def.(*cp.StateDefContext))
		}
	}
	panic("No state definition found!")
}

func (v *Visitor) visitState(ctx *cp.StateContext) models.Puzzle {
	if ctx.RANDOM() != nil {
		panic("random not implemented yet!")
	}
	if ctx.AdvancedState() != nil {
		panic("advanced state not implementd yet!")
	}
	visitor := newBeginnerStateVisitor()
	return visitor.visitBeginnerState(ctx.BeginnerState().(*cp.BeginnerStateContext))
}

func (v *Visitor) visitStateDef(ctx *cp.StateDefContext) models.Puzzle {
	return v.visitState(ctx.State().(*cp.StateContext))
}
