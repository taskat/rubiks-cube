package errorvisitor

import (
	"fmt"
	"strconv"

	ap "github.com/taskat/rubiks-cube/src/algo/parser"
	"github.com/taskat/rubiks-cube/src/basevisitor"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/symboltable"
)

type algorithmVisitor struct {
	basevisitor.ErrorVisitor
	turns *symboltable.Table[eh.IContext]
}

func newAlgorithmVisitor(ev basevisitor.ErrorVisitor, turns *symboltable.Table[eh.IContext]) *algorithmVisitor {
	return &algorithmVisitor{ErrorVisitor: ev, turns: turns}
}

func (v *algorithmVisitor) visitAlgorithm(ctx *ap.AlgorithmContext) {
	for _, turn := range ctx.AllTurn() {
		v.visitTurn(turn.(*ap.TurnContext))
	}
}

func (v *algorithmVisitor) visitTurn(ctx *ap.TurnContext) {
	fmt.Println("a" + ctx.GetText())
	if ctx.WORD() != nil {
		turnName := ctx.WORD().GetText()
		if ctx.PRIME() != nil {
			turnName += "'"
		} else if ctx.NUMBER() != nil {
			numberString := ctx.NUMBER().GetText()
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			if number == 2 {
				turnName += "2"
			} else {
				v.Eh().AddError(ctx, "Only 2 can be in turns", v.FileName())
			}
		}
		if _, err := v.turns.GetIdentifier(turnName); err != nil {
			v.Eh().AddError(ctx, "Turn "+turnName+" is not defined", v.FileName())
		}
	} else {
		v.visitAlgorithm(ctx.Algorithm().(*ap.AlgorithmContext))
	}
}
