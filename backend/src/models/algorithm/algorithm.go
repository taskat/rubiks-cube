package algorithm

type Algorithm Block

func NewAlgorithm(b Block) Algorithm {
	return b
}

type Setter func(block Block)
