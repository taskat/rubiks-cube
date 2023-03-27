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

type recognizer interface {
	AddErrorListener(antlr.ErrorListener)
	RemoveErrorListeners()
}

func addCustomErrorListener(to recognizer, fileName string) {
	to.RemoveErrorListeners()
	to.AddErrorListener(errorlistener.NewErrorCollector(fileName))
}

func readConfig(fileName string) {
	input := readFile(fileName)
	stream := getTokens(input, fileName)
	tree := buildTree(stream, fileName)
	checkTree(tree, fileName)
}

func readFile(fileName string) *antlr.FileStream {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		panic(err)
	}
	return input
}

func getTokens(input *antlr.FileStream, fileName string) *antlr.CommonTokenStream {
	lexer := configlexer.NewConfigLexer(input)
	addCustomErrorListener(lexer, fileName)
	return antlr.NewCommonTokenStream(lexer, 0)
}

func buildTree(stream *antlr.CommonTokenStream, fileName string) antlr.ParseTree {
	parser := configparser.NewConfigParser(stream)
	addCustomErrorListener(parser, fileName)
	parser.BuildParseTrees = true
	return parser.ConfigFile()
}

func checkTree(tree antlr.ParseTree, fileName string) {
	visitor := errorvisitor.NewVisitor(fileName)
	visitor.Visit(tree)
}

func printErrors() {
	allMessages := eh.GetAllMessages()
	fmt.Println("======")
	for level, msgs := range allMessages {
		if len(msgs) == 0 {
			fmt.Println("There were no " + level + "s")
		} else {
			fmt.Println(level + "s:")
			for _, msg := range msgs {
				fmt.Println("  " + msg.String())
			}
		}
	}
}

func main() {
	fileName := os.Args[1]
	readConfig(fileName)
	printErrors()
}
