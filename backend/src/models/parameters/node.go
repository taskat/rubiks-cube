package parameters

type Node struct {
	Sides []Side
	Index int
}

func NewNode(sides []Side, index int) *Node {
	return &Node{Sides: sides, Index: index}
}
