package command

import (
	"go-kafka-messaging/internal/app/receiver/client"
	"log"
)

func closeConsumer(input interface{}) {
	if nil == input {
		log.Println("Invalid input")
		return
	}

	message := input.(*client.KafkaTopic)
	log.Printf("Sutting down consumer %s \n", message.Name)
	message.IsRunning = false
}
