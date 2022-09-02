package inputparser

type Parser interface {
	Parse(text string) (*Message, error)
}

const (
	ACTION_COMMAND = "COMMAND"
	ACTION_MESSAGE = "MESSAGE"
)

var ACTIONS = [2]string{"EXIT", "HISTORY"}

type Message struct {
	Type     string
	Receiver string
	Value    string
}
