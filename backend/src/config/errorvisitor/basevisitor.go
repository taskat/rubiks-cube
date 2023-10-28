package errorvisitor

import (
	"github.com/taskat/rubiks-cube/src/basevisitor"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type baseVisitor struct {
	basevisitor.ErrorVisitor
	finished bool
	valid    bool
	colors   map[string]int
}

func newBaseVisitor(fileName string, errorHandler *eh.Errorhandler) *baseVisitor {
	return &baseVisitor{ErrorVisitor: *basevisitor.NewErrorVisitor(errorHandler, fileName), finished: true, valid: true,
		colors: make(map[string]int)}
}

func (v *baseVisitor) getColors() map[string]int {
	return v.colors
}

func (v *baseVisitor) setFinished(finished bool) {
	v.finished = finished
}

func a() {}
