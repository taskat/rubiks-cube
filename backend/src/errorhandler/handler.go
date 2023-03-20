package errorhandler

import "fmt"

type IMessage interface {
	GetLevel() string
	fmt.Stringer
}

type errorhandler struct {
	messages []IMessage
}

func newHandler() errorhandler {
	return errorhandler{messages: make([]IMessage, 0)}
}

func (e *errorhandler) addMessage(err Message) {
	e.messages = append(e.messages, err)
}

var handler = newHandler()

func AddError(ctx IContext, text, file string) {
	error := NewMessage(ctx, text, file, "ERROR")
	handler.addMessage(error)
}

func AddInfo(ctx IContext, text, file string) {
	info := NewMessage(ctx, text, file, "INFO")
	handler.addMessage(info)
}

func AddWarning(ctx IContext, text, file string) {
	warning := NewMessage(ctx, text, file, "Warning")
	handler.addMessage(warning)
}

func AddMessage(m Message) {
	handler.addMessage(m)
}

func GetMessages() []IMessage {
	return handler.messages
}
