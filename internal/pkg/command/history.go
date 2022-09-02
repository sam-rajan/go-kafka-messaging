package command

type HistoryCommand struct {
	data string
}

func NewHistoryCommand(message string) Command {
	return &TextCommand{data: message}
}

func (self *HistoryCommand) Execute() {
}
