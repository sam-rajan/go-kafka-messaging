package main

import (
	"bufio"
	"go-kafka-messaging/cmd/sender-app/app"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	messageSender := app.Init()

	topic := "first_topic"
	scanner := bufio.NewScanner(os.Stdin)
	var i int = 0
	for scanner.Scan() {
		key := "key-" + strconv.Itoa(i)
		value := scanner.Text()

		message := kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(value),
		}
		messageSender.Send(message)
	}
}
