// Code generated from grammars/ConfigParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package configparser // ConfigParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ConfigParser.
type ConfigParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConfigParser#configFile.
	VisitConfigFile(ctx *ConfigFileContext) interface{}

	// Visit a parse tree produced by ConfigParser#puzzleTypeDef.
	VisitPuzzleTypeDef(ctx *PuzzleTypeDefContext) interface{}

	// Visit a parse tree produced by ConfigParser#puzzleType.
	VisitPuzzleType(ctx *PuzzleTypeContext) interface{}
}
