package errorhandler

type Position struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

func newPosition(line, column int) Position {
	return Position{Line: line, Column: column}
}
