package basevisitor

import eh "github.com/taskat/rubiks-cube/src/errorhandler"

type ErrorVisitor struct {
	eh       *eh.Errorhandler
	filename string
}

func NewErrorVisitor(eh *eh.Errorhandler, filename string) *ErrorVisitor {
	return &ErrorVisitor{eh: eh, filename: filename}
}

func (v *ErrorVisitor) Eh() *eh.Errorhandler {
	return v.eh
}

func (v *ErrorVisitor) FileName() string {
	return v.filename
}
