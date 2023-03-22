package errorlistener

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/taskat/rubiks-cube/src/errorhandler"
)

type ErrorCollector struct {
	antlr.ErrorListener
	file string
}

func NewErrorCollector(file string) *ErrorCollector {
	return &ErrorCollector{file: file}
}

func (ec *ErrorCollector) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol interface{}, line int, column int,
	msg string, e antlr.RecognitionException) {
	ctx := errorhandler.NewContext(line, column)
	errorhandler.AddError(ctx, msg, ec.file)
}
