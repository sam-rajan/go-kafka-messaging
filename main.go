package main

func main() {
	configFile := "go-kafka-messaging.properties"
	kafkaProperties := KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	kafkaProducer := KafkaProducer{}
	kafkaProducer.Produce(kafkaProperties.Value)
	kafkaConsumer := KafkaConsumer{}
	kafkaConsumer.Consume(kafkaProperties.Value)

}
