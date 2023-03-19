package errorlistener

import "github.com/taskat/rubiks-cube/src/errorhandler"

type context struct {
	t token
}

func newContext(line, column int) errorhandler.Context {
	return &context{t: *newToken(line, column)}
}

func (c context) GetStart() errorhandler.Token {
	return &c.t
}

type token struct {
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
