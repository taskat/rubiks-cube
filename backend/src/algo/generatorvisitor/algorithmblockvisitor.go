package generatorvisitor

import (
	"fmt"
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/models/algorithm"
	"github.com/taskat/rubiks-cube/src/symboltable"
)

type algorithmBlockVisitor struct {
	turns *symboltable.Table[[]string]
}

func newAlgorithmBlockVisitor(turns *symboltable.Table[[]string]) *algorithmBlockVisitor {
	return &algorithmBlockVisitor{turns: turns}
}

func (v *algorithmBlockVisitor) visitAlgorithm(ctx *ap.AlgorithmContext) *algorithm.Action {
	moves := v.visitAlgorithmMoves(ctx)
	return algorithm.NewAction(moves)
}

func (v *algorithmBlockVisitor) visitAlgorithmMoves(ctx *ap.AlgorithmContext) []string {
	moves := make([]string, 0)
	for _, turn := range ctx.AllTurn() {
		moves = append(moves, v.visitTurn(turn.(*ap.TurnContext))...)
	}
	return moves
}

func (v *algorithmBlockVisitor) visitTurn(ctx *ap.TurnContext) []string {
	if ctx.WORD() != nil {
		name := ctx.GetText()
		moves, err := v.turns.GetIdentifier(name)
		if err != nil {
			panic(fmt.Errorf("Turn %s not found", name))
		}
		return moves
	}
	repeat, _ := strconv.Atoi(ctx.NUMBER().GetText())
	moves := v.visitAlgorithmMoves(ctx.Algorithm().(*ap.AlgorithmContext))
	expanded := make([]string, len(moves)*repeat)
	for i := 0; i < repeat; i++ {
		copy(expanded[i*len(moves):], moves)
	}
	return expanded
}
