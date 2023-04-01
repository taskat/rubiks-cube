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
	pos   Position
	level Level
}

func NewMessage(ctx IContext, text, file string, level Level) Message {
	return Message{text: text, file: file, pos: getPosition(ctx), level: level}
}

func (m Message) GetLevel() Level {
	return m.level
}

func (m Message) MarshalJSON() ([]byte, error) {
	return NewMessageJson(&m).marshal()
}

func (m Message) String() string {
	return fmt.Sprintf("%s in %s at line %d, column %d", m.text, m.file,
		m.pos.Line, m.pos.Column+1)
}

func getPosition(ctx IContext) Position {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	return Position{Line: line, Column: column}
}
