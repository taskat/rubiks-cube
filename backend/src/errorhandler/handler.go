package errorhandler

import "fmt"

type IMessage interface {
	GetLevel() Level
	fmt.Stringer
}

type errorhandler struct {
	messages map[Level][]IMessage
}

func newHandler() errorhandler {
	return errorhandler{messages: make(map[Level][]IMessage, 0)}
}

func (e *errorhandler) addMessage(m Message) {
	e.messages[m.GetLevel()] = append(e.messages[m.GetLevel()], m)
}

var handler = newHandler()

func AddError(ctx IContext, text, file string) {
	error := NewMessage(ctx, text, file, ERROR)
	handler.addMessage(error)
}

func AddInfo(ctx IContext, text, file string) {
	info := NewMessage(ctx, text, file, INFO)
	handler.addMessage(info)
}

func AddWarning(ctx IContext, text, file string) {
	warning := NewMessage(ctx, text, file, WARNING)
	handler.addMessage(warning)
}

func AddMessage(m Message) {
	handler.addMessage(m)
}

func GetAllMessages() []IMessage {
	errors := GetErrors()
	infos := GetInfos()
	warnings := GetWarnings()
	return append(errors, append(infos, warnings...)...)
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
