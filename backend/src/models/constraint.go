package models

type Constraint struct {
	Turns  []string
	Sides  []string
	Colors []string
	Size   int
}

func NewConstraint(turns, sides, colors []string, size int) *Constraint {
	return &Constraint{Turns: turns, Sides: sides, Colors: colors, Size: size}
}
