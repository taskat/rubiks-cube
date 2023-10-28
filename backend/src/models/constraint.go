package models

type Constraint struct {
	Turns  []string
	Sides  []string
	Colors []string
}

func NewConstraint(turns, sides, colors []string) *Constraint {
	return &Constraint{Turns: turns, Sides: sides, Colors: colors}
}
