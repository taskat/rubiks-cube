package parameters

type Node struct {
	Sides []string
}

func NewNode(sides []string) *Node {
	return &Node{Sides: sides}
}
