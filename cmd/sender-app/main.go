package main

import (
	"go-kafka-messaging/cmd/sender-app/app"
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
	messageSender := app.InitializeSender(configMap)

}
