package generatorvisitor

import (
	"fmt"
	"strconv"
	"strings"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/symboltable"
)

type algorithmVisitor struct {
	turns *symboltable.Table[[]string]
}

func newAlgorithmVisitor(turns *symboltable.Table[[]string]) *algorithmVisitor {
	return &algorithmVisitor{turns: turns}
}

func (v *algorithmVisitor) visitAlgorithm(ctx *ap.AlgorithmContext) []string {
	moves := make([]string, 0)
	for _, turn := range ctx.AllTurn() {
		moves = append(moves, v.visitTurn(turn.(*ap.TurnContext))...)
	}
	return moves
}

func (v *algorithmVisitor) visitTurn(ctx *ap.TurnContext) []string {
	if ctx.WORD() != nil {
		name := ctx.WORD().GetText()
		moves, err := v.turns.GetIdentifier(name)
		if err != nil {
			panic(fmt.Errorf("Turn %s not found", name))
		}
		if len(moves) > 1 {
			return moves
		}
		name = strings.ReplaceAll(name, "(", "")
		name = strings.ReplaceAll(name, ")", "")
		return []string{ctx.GetText()}
	}
	repeat, _ := strconv.Atoi(ctx.NUMBER(0).GetText())
	moves := v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
	expanded := make([]string, len(moves)*repeat)
	for i := 0; i < repeat; i++ {
		copy(expanded[i*len(moves):], moves)
	}
	return expanded
}
