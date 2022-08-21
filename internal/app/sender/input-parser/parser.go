package inputparser

import (
	"errors"
	"strings"
)

type Message struct {
	topic   string
	message string
}

func (self *Message) GetTopic() string {
	return self.topic
}

func (self *Message) GetMessage() string {
	return self.message
}

func ParseMessage(input string) (*Message, error) {
	if string(input[0]) != "@" {
		return new(Message), errors.New("Expecting a recepient reference")
	}

	receiverEnd := strings.Index(input, ":")

	if receiverEnd <= 1 {
		return new(Message), errors.New("Invalid Receiver reference")
	}

	receiverName := input[1:receiverEnd]
	parsedMessage := input[receiverEnd+1:]

	return &Message{topic: receiverName, message: strings.Trim(parsedMessage, " ")}, nil
}
