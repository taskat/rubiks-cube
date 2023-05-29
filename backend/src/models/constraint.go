package models

type Constraint struct {
	Turns []string
}

func NewConstraint(turns []string) *Constraint {
	return &Constraint{Turns: turns}
}
