package inputparser

const (
	ACTION_MESSAGE = "TEXT"
)

var ACTIONS = [2]string{"EXIT", "HISTORY"}

type Message struct {
	Type     string `json:"Type"`
	Receiver string `json:"Receiver"`
	Value    string `json:"Value"`
}
