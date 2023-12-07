package generatorvisitor

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	cp "github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/models"
	"github.com/taskat/rubiks-cube/src/models/cube"
	"github.com/taskat/rubiks-cube/src/models/parameters"
)

type Visitor struct {
	size            int
	puzzle          models.Puzzle
	sideConstructor func(string) parameters.Side
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
		v.visitConfigLine(line.(*cp.ConfigLineContext))
	}
	return v.puzzle
}

func (v *Visitor) visitConfigLine(ctx *cp.ConfigLineContext) {
	if ctx.PuzzleTypeDef() != nil {
		v.visitPuzzleTypeDef(ctx.PuzzleTypeDef().(*cp.PuzzleTypeDefContext))
	}
	if ctx.SizeDef() != nil {
		v.visitSizeDef(ctx.SizeDef().(*cp.SizeDefContext))
	}
	if ctx.StateDef() != nil {
		v.puzzle = v.visitStateDef(ctx.StateDef().(*cp.StateDefContext))
	}
}

func (v *Visitor) visitPuzzleTypeDef(ctx *cp.PuzzleTypeDefContext) {
	if ctx.PuzzleType().CUBE() != nil {
		puzzle := cube.NewCube(v.size)
		v.sideConstructor = puzzle.SideConstructor()
	}
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
		visitor := newAdvancedStateVisitor(v.size, v.sideConstructor)
		return visitor.visitAdvancedState(ctx.AdvancedState().(*cp.AdvancedStateContext))
	}
	visitor := newBeginnerStateVisitor(v.size, v.sideConstructor)
	return visitor.visitBeginnerState(ctx.BeginnerState().(*cp.BeginnerStateContext))
}

func (v *Visitor) visitStateDef(ctx *cp.StateDefContext) models.Puzzle {
	return v.visitState(ctx.State().(*cp.StateContext))
}
