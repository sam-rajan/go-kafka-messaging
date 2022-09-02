package command

type Factory interface {
	GetCommand(action string) (Command, error)
}

type CommandFactory struct {
}

func (self *CommandFactory) GetCommand(action string) (Command, error) {
	switch action {
	case "TEXT":
		return NewMessageCommand(""), nil
	case "HISTORY":
		return NewHistoryCommand(""), nil
	case "EXIT":
		return NewExitCommand(""), nil
	}

	return nil, nil
}
