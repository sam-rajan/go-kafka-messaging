package command

import "log"

type TextCommand struct {
	data string
}

func NewMessageCommand(message string) Command {
	return &TextCommand{data: message}
}

func (self *TextCommand) Execute() {
	if self.data == "" {
		return
	}

	log.Println(self.data)
}
