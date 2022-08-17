package app

import (
	"go-kafka-messaging/internal/app/receiver"
	"go-kafka-messaging/internal/app/receiver/client"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
)

func Init(kafkaProperties configreader.KafkaProperties) *receiver.MessageReceiver {
	kafkaClient := client.KafkaClient{}
	kafkaClient.Init(&kafkaProperties)

	return receiver.NewMessageReceiver(&kafkaClient)
}
