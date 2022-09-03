package inputparser

import (
	"errors"
	"strings"
)

func Parse(input string) (*Message, error) {

	receiver, err := "", error(nil)
	if string(input[0]) == "@" {
		receiver, input, err = parseReceiver(input)
	}

	if err != nil {
		return new(Message), err
	}

	data := &Message{}

	data.Type = ACTION_MESSAGE
	if isAction(input) {
		data.Type = input
	}

	if receiver == "" && data.Type == ACTION_MESSAGE {
		return new(Message), errors.New("Message without recepient is not possible")
	}

	data.Receiver = receiver
	data.Value = input

	return data, nil
}

func parseReceiver(input string) (string, string, error) {
	receiverEnd := strings.Index(input, " ")

	if receiverEnd == -1 {
		return "", input, errors.New("Expecting a recepient reference")
	}

	return input[1:receiverEnd], input[receiverEnd+1:], nil
}

func isAction(input string) bool {

	for _, v := range ACTIONS {
		if v == strings.ToUpper(input) {
			return true
		}
	}

	return false
}
