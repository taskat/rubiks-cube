package errorvisitor

import eh "github.com/taskat/rubiks-cube/src/errorhandler"

type baseVisitor struct {
	fileName string
	finished bool
	valid    bool
	eh       *eh.Errorhandler
	colors   map[string]int
}

func newBaseVisitor(fileName string, errorHandler *eh.Errorhandler) *baseVisitor {
	return &baseVisitor{fileName: fileName, finished: true, valid: true, eh: errorHandler,
		colors: make(map[string]int)}
}

func (v *baseVisitor) getColors() map[string]int {
	return v.colors
}

func (v *baseVisitor) getEh() *eh.Errorhandler {
	return v.eh
}

func (v *baseVisitor) getFileName() string {
	return v.fileName
}

func (v *baseVisitor) setFinished(finished bool) {
	v.finished = finished
}

func a() {}
