package errorlistener

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Recognizer interface {
	AddErrorListener(antlr.ErrorListener)
	RemoveErrorListeners()
}

func AddCustomErrorListener(to Recognizer, fileName string) {
	to.RemoveErrorListeners()
	to.AddErrorListener(NewErrorCollector(fileName))
}
