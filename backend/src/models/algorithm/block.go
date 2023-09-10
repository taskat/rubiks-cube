package algorithm

import "github.com/taskat/rubiks-cube/src/models"

type Block interface {
	Execute(p models.Puzzle) Block
	Finished() bool
}
