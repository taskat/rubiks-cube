package errorhandler

import "encoding/json"

type MessageJSON struct {
	Text  string   `json:"text"`
	File  string   `json:"file"`
	Pos   Position `json:"pos"`
	Level string   `json:"level"`
}

func NewMessageJson(m *Message) MessageJSON {
	return MessageJSON{
		Text:  m.text,
		File:  m.file,
		Pos:   m.pos,
		Level: m.level.String(),
	}
}

func (m MessageJSON) marshal() ([]byte, error) {
	return json.Marshal(m)
}
