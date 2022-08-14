package main

import (
	receiver "go-kafka-messaging/internal/app/receiver"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configFile := "kafka-consumer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	var wg = &sync.WaitGroup{}
	channel := make(chan kafka.Message, 5)

	kafkaConsumer := receiver.KafkaConsumer{}
	kafkaConsumer.Init(kafkaProperties.Value)

	messageReceiver := receiver.MessageReceiver{}
	messageReceiver.KafkaConsumer = kafkaConsumer.Instance

	topic := "first_topic"
	processor := receiver.KafkaProcessor{}
	processor.Process(wg, channel)
	messageReceiver.StartReceive(topic, wg, channel)

	wg.Wait()
}
