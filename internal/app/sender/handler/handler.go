package handler

import (
	"encoding/json"
	"go-kafka-messaging/internal/app/sender"
	inputparser "go-kafka-messaging/internal/app/sender/input-parser"
	"log"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type InputHandler struct {
	messageCount  int
	messageSender sender.MessageSender
	parser        inputparser.Parser
}

func NewMessageHandler(sender sender.MessageSender) sender.ReaderListener {
	return &InputHandler{messageSender: sender}
}

func (self *InputHandler) OnInputRead(message string) {
	parsedMessage, err := (self.parser).Parse(message)

	if err != nil {
		log.Printf("Input Parsing failed. Reason : %s\n", err.Error())
		return
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
}
