package handler

import (
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	"log"
)

func OnInputRead(message string) {

	parsedMessage, err := inputparser.Parse(message)

	if err != nil {
		log.Printf("Input Parsing failed. Reason : %s\n", err.Error())
		return
	}

	var command func(interface{})
	if parsedMessage.Receiver == "" {
		command = GetCommand(parsedMessage.Type)
	} else {
		command = sendMessage
	}

	command(*parsedMessage)
}
