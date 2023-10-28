package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Finish struct{}

func NewFinish() *Finish {
	return &Finish{}
}

func (f *Finish) Execute(p models.Puzzle) Block {
	return f
}

func (f *Finish) Finished() bool {
	return true
}
