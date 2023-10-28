package parameters

type Piece struct {
	Node
}

func NewPiece(n Node) *Piece {
	return &Piece{Node: n}
}
