package sender

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageSender struct {
	KafkaProducer *kafka.Producer
}

func (self *MessageSender) Send(message kafka.Message) {
	fmt.Println("Sending Message")
	self.KafkaProducer.Produce(&message, nil)
}
