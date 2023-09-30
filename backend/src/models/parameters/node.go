package parameters

type Node struct {
	Sides []Side
}

func NewNode(sides []string) *Node {
	typedSides := make([]Side, len(sides))
	for i, side := range sides {
		typedSides[i] = NewSide(side)
	}
	return &Node{Sides: typedSides}
}
