package errorhandler

import (
	"fmt"
)

type Context interface {
	GetStart() Token
}

type Token interface {
	GetLine() int
	GetColumn() int
}

type Error struct {
	text string
	file string
	pos  position
}

func NewError(ctx Context, text, file string) Error {
	return Error{text: text, file: file, pos: getPosition(ctx)}
}

func (e Error) String() string {
	return fmt.Sprintf("in %s, %s at line %d, column %d", e.file, e.text,
		e.pos.line, e.pos.column)
}

func getPosition(ctx Context) position {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	return position{line: line, column: column}
}
