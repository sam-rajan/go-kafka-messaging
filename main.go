package main

import "sync"

func main() {
	configFile := "go-kafka-messaging.properties"
	kafkaProperties := KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	var wg = &sync.WaitGroup{}
	kafkaProducer := KafkaProducer{}
	kafkaProducer.Produce(kafkaProperties.Value, wg)
	kafkaConsumer := KafkaConsumer{}
	kafkaConsumer.Consume(kafkaProperties.Value, wg)

	wg.Wait()
}
