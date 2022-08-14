package main

import (
	sender "go-kafka-messaging/internal/app/sender"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	configFile := "configs/kafka-producer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	kafkaProducer := sender.KafkaProducer{}
	kafkaProducer.Init(kafkaProperties.Value)

	messageSender := sender.MessageSender{}
	messageSender.KafkaProducer = kafkaProducer.Instance

	topic := "first_topic"
	for i := 0; i < 10; i++ {
		key := "key-" + strconv.Itoa(i)
		value := "value-" + time.Now().String()

		message := kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(value),
		}

		messageSender.Send(message)
		time.Sleep(10 * time.Second)
	}
}
