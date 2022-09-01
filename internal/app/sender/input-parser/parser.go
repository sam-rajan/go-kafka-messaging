package inputparser

type Parser interface {
	Parse(text string)
}

const (
	ACTION_COMMAND = "COMMAND"
	ACTION_MESSAGE = "MESSAGE"
)

type Message struct {
	Type     string
	Receiver string
	Value    string
}
