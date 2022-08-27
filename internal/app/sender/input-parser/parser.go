package inputparser

import (
	"errors"
	"strings"
)

type Action string

type Message struct {
	topic   string
	payload *Payload
}

type Payload struct {
	Action  Action
	Message string
}

func (self *Message) GetTopic() string {
	return self.topic
}

func (self *Message) GetPayload() *Payload {
	return self.payload
}

func ParseMessage(input string) (*Message, error) {

	receiver, err := "", error(nil)
	if string(input[0]) == "@" {
		receiver, input, err = parseReceiver(input)
	}

	if err != nil {
		return new(Message), errors.New(err.Error())
	}

	data := &Payload{}
	if isAction(input) {
		data.Action = Action(input)
	} else {
		data.Message = input
	}

	if receiver == "" && data.Message != "" {
		return new(Message), errors.New("Message without recepient is not possible")
	}

	return &Message{topic: receiver, payload: data}, nil
}

func parseReceiver(input string) (string, string, error) {
	receiverEnd := strings.Index(input, " ")

	if receiverEnd == -1 {
		return "", input, errors.New("Expecting a recepient reference")
	}

	return input[1:receiverEnd], input[receiverEnd+1:], nil
}

func isAction(input string) bool {

	actionList := []string{"EXIT", "HISTORY"}
	for _, v := range actionList {
		if v == strings.ToUpper(input) {
			return true
		}
	}

	return false
}
