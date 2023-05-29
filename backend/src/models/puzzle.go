package models

import (
	"encoding/json"
	"fmt"
)

type Puzzle interface {
	json.Marshaler
	fmt.Stringer
	GetValidator() Validator
	GetConstraint() Constraint
}

type Validator interface {
	Validate() []string
}
