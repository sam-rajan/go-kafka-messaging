package app

import (
	"go-kafka-messaging/internal/app/sender"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
)

func Init() *sender.MessageSender {
	configFile := "configs/kafka-producer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	kafkaProducer := sender.KafkaProducer{}
	kafkaProducer.Init(kafkaProperties.Value)

	messageSender := sender.MessageSender{}
	messageSender.KafkaProducer = kafkaProducer.Instance

	return &messageSender

}
