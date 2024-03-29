package errorhandler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Context struct {
	t token
}

func NewContext(line, column int) IContext {
	return &Context{t: *newToken(line, column)}
}

func NewContextFromTerminal(t antlr.TerminalNode) IContext {
	return &Context{t: *newToken(t.GetSymbol().GetLine(), t.GetSymbol().GetColumn())}
}

func (c Context) GetStart() antlr.Token {
	return &c.t
}

type token struct {
	antlr.Token  // only to suppress error, only GetLine and GetColumn are used
	line, column int
}

func newToken(line, column int) *token {
	return &token{line: line, column: column}
}

func (t token) GetLine() int {
	return t.line
}

func (t token) GetColumn() int {
	return t.column
}
