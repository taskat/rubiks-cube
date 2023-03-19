// Code generated from grammars/ConfigParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package configparser // ConfigParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConfigParserListener is a complete listener for a parse tree produced by ConfigParser.
type BaseConfigParserListener struct{}

var _ ConfigParserListener = &BaseConfigParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConfigParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConfigParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConfigParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConfigParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConfigFile is called when production configFile is entered.
func (s *BaseConfigParserListener) EnterConfigFile(ctx *ConfigFileContext) {}

// ExitConfigFile is called when production configFile is exited.
func (s *BaseConfigParserListener) ExitConfigFile(ctx *ConfigFileContext) {}

// EnterPuzzleTypeDef is called when production puzzleTypeDef is entered.
func (s *BaseConfigParserListener) EnterPuzzleTypeDef(ctx *PuzzleTypeDefContext) {}

// ExitPuzzleTypeDef is called when production puzzleTypeDef is exited.
func (s *BaseConfigParserListener) ExitPuzzleTypeDef(ctx *PuzzleTypeDefContext) {}

// EnterPuzzleType is called when production puzzleType is entered.
func (s *BaseConfigParserListener) EnterPuzzleType(ctx *PuzzleTypeContext) {}

// ExitPuzzleType is called when production puzzleType is exited.
func (s *BaseConfigParserListener) ExitPuzzleType(ctx *PuzzleTypeContext) {}
