package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/taskat/rubiks-cube/src/config/lexer"
	"github.com/taskat/rubiks-cube/src/config/parser"
)

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}
	lexer := configlexer.NewConfigLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := configparser.NewConfigParser(stream)
	parser.BuildParseTrees = true
	tree := parser.ConfigFile()
	fmt.Println(tree.GetText())
}
