package errorlistener

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type ErrorCollector struct {
	antlr.DefaultErrorListener
	file string
	eh   *eh.Errorhandler
}

func NewErrorCollector(file string, errorHandler *eh.Errorhandler) *ErrorCollector {
	return &ErrorCollector{file: file, eh: errorHandler}
}

func (ec *ErrorCollector) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol interface{}, line int, column int,
	msg string, e antlr.RecognitionException) {
	ctx := eh.NewContext(line, column)
	ec.eh.AddError(ctx, msg, ec.file)
}
