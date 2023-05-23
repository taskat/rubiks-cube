package scope

import (
	"fmt"

	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type ErrorScope struct {
	identifiers map[string][]eh.IContext
}

func NewErrorScope() *ErrorScope {
	return &ErrorScope{identifiers: make(map[string][]eh.IContext)}
}

func (s *ErrorScope) AddIdentifier(name string, ctx eh.IContext) {
	found := s.identifiers[name]
	found = append(found, ctx)
	s.identifiers[name] = found
}

func (s *ErrorScope) CheckForErrors(eh *eh.Errorhandler, filename string) {
	for name, found := range s.identifiers {
		if len(found) > 1 {
			msg := fmt.Sprintf("Identifier %s is defined multiple times", name)
			for _, ctx := range found {
				eh.AddError(ctx, msg, filename)
			}
		}
	}
}

func (s *ErrorScope) GetIdentifier(name string) (eh.IContext, error) {
	found := s.identifiers[name]
	if len(found) == 0 {
		return nil, fmt.Errorf("Identifier %s not found", name)
	}
	return found[0], nil
}
