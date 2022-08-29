package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-kafka-messaging/cmd/sender-app/app"
	inputparser "go-kafka-messaging/internal/app/sender/input-parser"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	"log"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	configFile := "configs/kafka-producer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	if err := app.InitializeBroadcastTopic(kafkaProperties); err != nil {
		log.Fatalln("Failed to initialize app")
	}
	messageSender := app.InitializeSender(&kafkaProperties)

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

		if topic == "" {

		}
		jsonString, _ := json.Marshal(message.GetPayload())

		kafkaMessage := kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(jsonString),
		}
		messageSender.Send(kafkaMessage)
		i = i + 1
	}
}
