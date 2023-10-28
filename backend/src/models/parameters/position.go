package parameters

type Position struct {
	Node
}

func NewPosition(n Node) *Position {
	return &Position{Node: n}
}
