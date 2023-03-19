package errorhandler

type position struct {
	line, column int
}

func newPosition(line, column int) position {
	return position{line: line, column: column}
}
