package scope

import (
	"fmt"

	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Scope[T any] struct {
	identifiers map[string]T
}

func NewScope[T any]() *Scope[T] {
	return &Scope[T]{identifiers: make(map[string]T)}
}

func (s *Scope[T]) AddIdentifier(name string, value T) {
	s.identifiers[name] = value
}

func (s *Scope[T]) CheckForErrors(eh *eh.Errorhandler, filename string) {}

func (s *Scope[T]) GetIdentifier(name string) (T, error) {
	found, ok := s.identifiers[name]
	if !ok {
		var t T
		return t, fmt.Errorf("Identifier %s not found", name)
	}
	return found, nil
}
