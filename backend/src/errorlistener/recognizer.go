package errorlistener

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Recognizer interface {
	AddErrorListener(antlr.ErrorListener)
	RemoveErrorListeners()
}

func AddCustomErrorListener(to Recognizer, fileName string, eh eh.Errorhandler) {
	to.RemoveErrorListeners()
	to.AddErrorListener(NewErrorCollector(fileName, eh))
}
