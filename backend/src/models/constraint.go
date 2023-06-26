package models

type Constraint struct {
	Turns []string
	Sides []string
}

func NewConstraint(turns, sides []string) *Constraint {
	return &Constraint{Turns: turns, Sides: sides}
}
