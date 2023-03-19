// Code generated from grammars/ConfigParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package configparser // ConfigParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConfigParserListener is a complete listener for a parse tree produced by ConfigParser.
type ConfigParserListener interface {
	antlr.ParseTreeListener

	// EnterConfigFile is called when entering the configFile production.
	EnterConfigFile(c *ConfigFileContext)

	// EnterPuzzleTypeDef is called when entering the puzzleTypeDef production.
	EnterPuzzleTypeDef(c *PuzzleTypeDefContext)

	// EnterPuzzleType is called when entering the puzzleType production.
	EnterPuzzleType(c *PuzzleTypeContext)

	// ExitConfigFile is called when exiting the configFile production.
	ExitConfigFile(c *ConfigFileContext)

	// ExitPuzzleTypeDef is called when exiting the puzzleTypeDef production.
	ExitPuzzleTypeDef(c *PuzzleTypeDefContext)

	// ExitPuzzleType is called when exiting the puzzleType production.
	ExitPuzzleType(c *PuzzleTypeContext)
}
