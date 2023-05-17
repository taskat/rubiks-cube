package errorvisitor

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Visitor struct {
	fileName string
	eh       *eh.Errorhandler
}

func NewVisitor(fileName string, eh *eh.Errorhandler) *Visitor {
	return &Visitor{fileName: fileName, eh: eh}
}

func (v *Visitor) Visit(tree antlr.Tree) {
}
