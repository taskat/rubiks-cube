package errorhandler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type IContext interface {
	GetStart() antlr.Token
}

type Error struct {
	text string
	file string
	pos  position
}

func NewError(ctx IContext, text, file string) Error {
	return Error{text: text, file: file, pos: getPosition(ctx)}
}

func (e Error) String() string {
	return fmt.Sprintf("%s in %s at line %d, column %d", e.text, e.file,
		e.pos.line, e.pos.column+1)
}

func getPosition(ctx IContext) position {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	return position{line: line, column: column}
}
