package errorhandler

import "fmt"

type Errorhandler struct {
	messages map[Level][]Message
}

func NewHandler() Errorhandler {
	m := make(map[Level][]Message, 0)
	m["ERROR"] = make([]Message, 0)
	m["WARNING"] = make([]Message, 0)
	m["INFO"] = make([]Message, 0)
	return Errorhandler{messages: m}
}

func (e *Errorhandler) AddMessage(m Message) {
	e.messages[m.GetLevel()] = append(e.messages[m.GetLevel()], m)
}

func (e *Errorhandler) AddError(ctx IContext, text, file string) {
	error := NewMessage(ctx, text, file, ERROR)
	e.AddMessage(error)
}

func (e *Errorhandler) AddInfo(ctx IContext, text, file string) {
	info := NewMessage(ctx, text, file, INFO)
	e.AddMessage(info)
}

func (e *Errorhandler) AddWarning(ctx IContext, text, file string) {
	warning := NewMessage(ctx, text, file, WARNING)
	e.AddMessage(warning)
}

func (e *Errorhandler) GetMessages() []Message {
	var msgs []Message
	for _, m := range e.messages {
		msgs = append(msgs, m...)
	}
	return msgs
}

func (e *Errorhandler) GetMessagesByLevel() map[Level][]Message {
	return e.messages
}

func (e *Errorhandler) GetErrors() []Message {
	return e.messages["ERROR"]
}

func (e *Errorhandler) GetInfos() []Message {
	return e.messages["INFO"]
}

func (e *Errorhandler) GetWarnings() []Message {
	return e.messages["WARNING"]
}

func (e *Errorhandler) HasErrors() bool {
	return len(e.messages["ERROR"]) > 0
}

func (e *Errorhandler) PrintAll() {
	fmt.Println("======")
	for level, msgs := range e.messages {
		if len(msgs) == 0 {
			fmt.Println("There were no " + level + "s")
		} else {
			fmt.Println(level + "s:")
			for _, msg := range msgs {
				fmt.Println("  " + msg.String())
			}
		}
	}
}
