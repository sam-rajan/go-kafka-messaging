package command

import (
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	"log"
)

func printMessage(input interface{}) {

	if nil == input {
		log.Println("Invalid input")
		return
	}

	message := input.(inputparser.Message)

	log.Printf("Received message in %s consumer\n", message.Receiver)
	log.Printf("Received Message = %s\n ", string(message.Value))
}
