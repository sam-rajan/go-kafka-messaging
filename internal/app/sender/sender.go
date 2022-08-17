package sender

import (
	"fmt"
	"go-kafka-messaging/internal/app/sender/client"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageSender struct {
	kafkaClient *client.KafkaClient
}

func NewMessageSender(client *client.KafkaClient) *MessageSender {
	return &MessageSender{kafkaClient: client}
}

func (self *MessageSender) Send(message kafka.Message) {
	fmt.Println("Sending Message")

	producer := self.kafkaClient.Get()

	if nil == producer {
		panic("No Kafka producer instance found")
	}

	producer.Produce(&message, nil)
}
