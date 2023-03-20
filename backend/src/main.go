package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/taskat/rubiks-cube/src/config/errorvisitor"
	"github.com/taskat/rubiks-cube/src/config/lexer"
	"github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/errorlistener"
)

func main() {
	fileName := os.Args[1]
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		panic(err)
	}
	lexer := configlexer.NewConfigLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorlistener.NewErrorCollector(fileName))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := configparser.NewConfigParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorlistener.NewErrorCollector(fileName))
	parser.BuildParseTrees = true
	tree := parser.ConfigFile()
	visitor := errorvisitor.NewVisitor(fileName)
	visitor.Visit(tree)

	fmt.Println("======")
	errors := eh.GetErrors()
	if len(errors) == 0 {
		fmt.Println("There were no errors")
	} else {
		fmt.Println("Errors:")
		for _, err := range errors {
			fmt.Println("  " + err.String())
		}
	}
	warnings := eh.GetWarnings()
	if len(warnings) == 0 {
		fmt.Println("There were no warnings")
	} else {
		fmt.Println("Warnings:")
		for _, warning := range warnings {
			fmt.Println("  " + warning.String())
		}
	}
	infos := eh.GetInfos()
	if len(infos) == 0 {
		fmt.Println("There were no infos")
	} else {
		fmt.Println("Infos:")
		for _, info := range infos {
			fmt.Println("  " + info.String())
		}
	}

}
