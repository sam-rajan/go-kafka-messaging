package app

import (
	"go-kafka-messaging/internal/app/sender"
	"go-kafka-messaging/internal/app/sender/client"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
)

func Init() *sender.MessageSender {
	configFile := "configs/kafka-producer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	kafkaClient := client.KafkaClient{}
	kafkaClient.Init(&kafkaProperties)

	messageSender := sender.NewMessageSender(&kafkaClient)
	return messageSender
}
