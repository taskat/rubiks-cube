package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/taskat/rubiks-cube/src/config/errorvisitor"
	"github.com/taskat/rubiks-cube/src/config/generatorvisitor"
	"github.com/taskat/rubiks-cube/src/config/lexer"
	"github.com/taskat/rubiks-cube/src/config/parser"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/errorlistener"
	"github.com/taskat/rubiks-cube/src/models"
)

type recognizer interface {
	AddErrorListener(antlr.ErrorListener)
	RemoveErrorListeners()
}

func addCustomErrorListener(to recognizer, fileName string) {
	to.RemoveErrorListeners()
	to.AddErrorListener(errorlistener.NewErrorCollector(fileName))
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

func createCube(tree antlr.Tree) models.Puzzle {
	visitor := generatorvisitor.NewVisitor()
	return visitor.Visit(tree)
}

func getTokens(input *antlr.FileStream, fileName string) *antlr.CommonTokenStream {
	lexer := configlexer.NewConfigLexer(input)
	addCustomErrorListener(lexer, fileName)
	return antlr.NewCommonTokenStream(lexer, 0)
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

func readConfig(fileName string) antlr.Tree {
	input := readFile(fileName)
	stream := getTokens(input, fileName)
	tree := buildTree(stream, fileName)
	checkTree(tree, fileName)
	return tree
}

func readFile(fileName string) *antlr.FileStream {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		panic(err)
	}
	return input
}

func main() {
	fileName := os.Args[1]
	tree := readConfig(fileName)
	if eh.HasErrors() {
		printErrors()
		os.Exit(1)
	}
	cube := createCube(tree)
	fmt.Println(cube)

	printErrors()
}
