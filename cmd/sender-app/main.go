package main

import (
	"bufio"
	"fmt"
	"go-kafka-messaging/cmd/sender-app/app"
	inputparser "go-kafka-messaging/internal/app/sender/input-parser"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	messageSender := app.Init()

	scanner := bufio.NewScanner(os.Stdin)
	var i int = 0
	for scanner.Scan() {

		value := scanner.Text()

		if value == "" {
			continue
		}

		message, err := inputparser.ParseMessage(value)
		if err != nil {
			fmt.Printf("Input Parsing failed. Reason : %s\n", err.Error())
			continue
		}

		key := "key-" + strconv.Itoa(i)
		topic := message.GetTopic()
		kafkaMessage := kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(message.GetMessage()),
		}
		messageSender.Send(kafkaMessage)
		i = i + 1
	}
}
