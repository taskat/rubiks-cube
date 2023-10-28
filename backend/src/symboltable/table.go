package symboltable

import (
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type scope[T any] interface {
	AddIdentifier(name string, value T)
	CheckForErrors(eh *eh.Errorhandler, filename string)
	GetIdentifier(name string) (T, error)
}

type Table[T any] struct {
	scopes []scope[T]
}

func NewTable[T any]() *Table[T] {
	return &Table[T]{scopes: make([]scope[T], 0)}
}

func (t *Table[T]) AddIdentifier(name string, value T) {
	t.scopes[len(t.scopes)-1].AddIdentifier(name, value)
}

func (t *Table[T]) CheckForErrorsTop(eh *eh.Errorhandler, filename string) {
	t.scopes[len(t.scopes)-1].CheckForErrors(eh, filename)
}

func (t *Table[T]) GetIdentifier(name string) (T, error) {
	var err error
	for i := len(t.scopes) - 1; i >= 0; i-- {
		var found T
		found, err = t.scopes[i].GetIdentifier(name)
		if err == nil {
			return found, nil
		}
	}
	var zeroValue T
	return zeroValue, err
}

func (t *Table[T]) PushScope(scope scope[T]) {
	t.scopes = append(t.scopes, scope)
}

func (t *Table[T]) PopScope() {
	t.scopes = t.scopes[:len(t.scopes)-1]
}
