package models

import (
	"encoding/json"
	"fmt"
)

type Puzzle interface {
	json.Marshaler
	fmt.Stringer
}
