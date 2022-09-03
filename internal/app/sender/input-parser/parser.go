package inputparser

const (
	ACTION_MESSAGE = "TEXT"
)

var ACTIONS = [2]string{"EXIT", "HISTORY"}

type Message struct {
	Type     string
	Receiver string
	Value    string
}
