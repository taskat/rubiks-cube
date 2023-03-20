package errorhandler

import "fmt"

type IMessage interface {
	GetLevel() string
	fmt.Stringer
}

type errorhandler struct {
	messages map[string][]IMessage
}

func newHandler() errorhandler {
	return errorhandler{messages: make(map[string][]IMessage, 0)}
}

func (e *errorhandler) addMessage(err Message) {
	e.messages[err.GetLevel()] = append(e.messages[err.GetLevel()], err)
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

func GetAllMessages() []IMessage {
	return nil
}

func GetErrors() []IMessage {
	return handler.messages["ERROR"]
}

func GetInfos() []IMessage {
	return handler.messages["INFO"]
}

func GetWarnings() []IMessage {
	return handler.messages["WARNING"]
}