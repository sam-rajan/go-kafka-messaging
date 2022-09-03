package handler

import (
	"encoding/json"
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	"log"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func sendMessage(input interface{}) {

	if nil == input {
		log.Println("Invalid input")
		return
	}

	message := input.(inputparser.Message)
	jsonString, _ := json.Marshal(message)

	key := strconv.Itoa(messageCounter)
	kafkaMessage := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &message.Receiver, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(jsonString),
	}
	messageSender.Send(kafkaMessage)

	messageCounter = messageCounter + 1
}
