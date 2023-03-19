// Code generated from grammars/ConfigParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package configparser // ConfigParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseConfigParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseConfigParserVisitor) VisitConfigFile(ctx *ConfigFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigParserVisitor) VisitPuzzleTypeDef(ctx *PuzzleTypeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConfigParserVisitor) VisitPuzzleType(ctx *PuzzleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
