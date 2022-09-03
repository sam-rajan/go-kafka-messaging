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

	command := GetCommand(parsedMessage.Type)

	if nil == command {
		log.Printf("Invalid command %s \n", parsedMessage.Value)
		return
	}
	command(*parsedMessage)
}
