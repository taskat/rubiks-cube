package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/taskat/rubiks-cube/src/config/errorvisitor"
	"github.com/taskat/rubiks-cube/src/config/lexer"
	"github.com/taskat/rubiks-cube/src/config/parser"
	"github.com/taskat/rubiks-cube/src/errorhandler"
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
	errors := errorhandler.GetErrors()
	if len(errors) == 0 {
		fmt.Println("No errors found")
	} else {
		for _, err := range errors {
			fmt.Println(err.String())
		}
	}
}
