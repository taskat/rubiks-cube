package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Label struct {
	name string
	next Block
}

func NewLabel(name string) *Label {
	return &Label{name: name, next: nil}
}

func (l *Label) Execute(p models.Puzzle) Block {
	p.(puzzleWrapper).StartStep(l.name)
	return l.next
}

func (l *Label) Finished() bool {
	return false
}

func (l *Label) NextSetter() Setter {
	return func(block Block) {
		l.next = block
	}
}

type puzzleWrapper interface {
	models.Puzzle
	StartStep(string)
}
