package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-kafka-messaging/internal/app/sender"
	inputparser "go-kafka-messaging/internal/app/sender/input-parser"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsoleMessageReader struct {
}

func NewMessageReader() sender.MessageReader {

}

func (self *ConsoleMessageReader) ReadMessage(listener sender.ReaderListener) {
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
