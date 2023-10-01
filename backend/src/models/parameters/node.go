package parameters

type Node struct {
	Sides []Side
}

func NewNode(sides []Side) *Node {
	return &Node{Sides: sides}
}
