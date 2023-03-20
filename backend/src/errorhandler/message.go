package errorhandler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type IContext interface {
	GetStart() antlr.Token
}

type Message struct {
	text  string
	file  string
	pos   position
	level string
}

func NewMessage(ctx IContext, text, file, level string) Message {
	return Message{text: text, file: file, pos: getPosition(ctx), level: level}
}

func (m Message) GetLevel() string {
	return m.level
}

func (m Message) String() string {
	return fmt.Sprintf("%s in %s at line %d, column %d", m.text, m.file,
		m.pos.line, m.pos.column+1)
}

func getPosition(ctx IContext) position {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	return position{line: line, column: column}
}
