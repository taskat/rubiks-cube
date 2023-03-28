package models

import "fmt"

type Puzzle interface {
	ToJSON() []byte
	fmt.Stringer
}
