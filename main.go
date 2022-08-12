package main

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configFile := "go-kafka-messaging.properties"
	kafkaProperties := KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	var wg = &sync.WaitGroup{}
	kafkaProducer := KafkaProducer{}
	kafkaProducer.Produce(kafkaProperties.Value, wg)

	channel := make(chan kafka.Message, 5)
	kafkaConsumer := KafkaConsumer{}
	kafkaConsumer.Consume(kafkaProperties.Value, wg, channel)
	kafkaProcessor := KafkaProcessor{}
	kafkaProcessor.Process(wg, channel)

	wg.Wait()
}
