package main

import (
	"go-kafka-messaging/cmd/sender-app/app"
	"go-kafka-messaging/internal/app/sender"
	"go-kafka-messaging/internal/app/sender/handler"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	"log"
)

func main() {

	configFile := "configs/kafka-producer.properties"
	configReader := configreader.NewFileConfig(configFile)
	configMap := configReader.ReadConfig()

	if err := app.InitializeBroadcastTopic(configMap); err != nil {
		log.Fatalln("Failed to initialize app")
	}
	messageSender := sender.NewMessageSender(configMap)

	inputListener := handler.NewMessageHandler(messageSender)
	inputReader := app.NewInputReader()
	inputReader.ReadMessage(inputListener)
}
