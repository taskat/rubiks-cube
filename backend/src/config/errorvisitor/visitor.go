package errorvisitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/errorhandler"
)

type Visitor struct {
	fileName string
}

func NewVisitor(fileName string) *Visitor {
	return &Visitor{fileName: fileName}
}

func (v *Visitor) Visit(tree antlr.Tree) {
	ctx, ok := tree.(*configparser.ConfigFileContext)
	if !ok {
		panic("Invalid tree")
	}
	v.visitConfigFile(ctx)
}

func (v *Visitor) visitConfigFile(ctx *configparser.ConfigFileContext) {
	puzzleDefs := getLines[*configparser.PuzzleTypeDefContext](ctx.AllConfigLine())
	checkOnlyOneDef(puzzleDefs, "puzzle definition", v.fileName)
	sizeDefs := getLines[*configparser.SizeDefContext](ctx.AllConfigLine())
	checkOnlyOneDef(sizeDefs, "size definition", v.fileName)
	stateDescriptionDefs := getLines[*configparser.StateDescriptionDefContext](ctx.AllConfigLine())
	checkOnlyOneDef(stateDescriptionDefs, "state description definition", v.fileName)
	stateDefs := getLines[*configparser.StateDefContext](ctx.AllConfigLine())
	checkOnlyOneDef(stateDefs, "state definition", v.fileName)
}

func (v *Visitor) visitPuzzleTypeDef(ctx *configparser.PuzzleTypeDefContext) {
}

func checkOnlyOneDef[def antlr.ParserRuleContext](defs []def, defType string, fileName string) {
	if len(defs) > 1 {
		for _, d := range defs {
			e := errorhandler.NewError(d, "Multiple "+defType+" found", fileName)
			errorhandler.AddError(e)
		}
	} else if len(defs) == 0 {
		ctx := errorhandler.NewContext(-1, -1)
		e := errorhandler.NewError(ctx, "No "+defType+" found", fileName)
		errorhandler.AddError(e)
	}
}

func getLines[def any](lines []configparser.IConfigLineContext) []def {
	result := make([]def, 0, 1)
	for _, l := range lines {
		line := l.(*configparser.ConfigLineContext).GetChild(0)
		if line, ok := line.(def); ok {
			result = append(result, line)
		}
	}
	return result
}
