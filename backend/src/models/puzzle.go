package models

import (
	"encoding/json"
	"fmt"
)

type Puzzle interface {
	json.Marshaler
	fmt.Stringer
	GetValidator() Validator
}

type Validator interface {
	Validate() []string
}
