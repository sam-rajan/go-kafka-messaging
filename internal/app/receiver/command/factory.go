package command

const (
	TEXT    = "TEXT"
	HISTORY = "HISTORY"
	EXIT    = "EXIT"
)

func GetCommand(action string) func(input interface{}) {

	switch action {
	default:
		return printMessage
	case HISTORY:
		return getHistory
	case EXIT:
		return closeConsumer
	}
}
