package handler

import (
	"encoding/json"
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	"log"
	"strconv"
)

func sendMessage(input interface{}) {

	if nil == input {
		log.Println("Invalid input")
		return
	}

	message := input.(inputparser.Message)
	jsonString, _ := json.Marshal(message)

	key := strconv.Itoa(messageCounter)
	messageSender.Send(message.Receiver, key, string(jsonString))

	messageCounter = messageCounter + 1
}
