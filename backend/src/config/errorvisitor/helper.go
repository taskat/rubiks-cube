package errorvisitor

import (
	"fmt"

	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type iVisitor interface {
	Eh() *eh.Errorhandler
	getColors() map[string]int
	FileName() string
	setFinished(bool)
}

func checkColors(errorCtx eh.IContext, visitor iVisitor, expected int) {
	colors := visitor.getColors()
	for color, amount := range colors {
		if amount > expected || amount < expected {
			warningMsg := fmt.Sprintf("color %s is defined %d times, should be %d times", color, amount, expected)
			visitor.Eh().AddWarning(errorCtx, warningMsg, visitor.FileName())
			visitor.setFinished(false)
		}
	}
}
