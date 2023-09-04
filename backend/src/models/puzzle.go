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
	Turn(name string)
}

type Validator interface {
	Validate() []string
}
