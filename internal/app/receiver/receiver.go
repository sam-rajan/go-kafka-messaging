package receiver

import (
	"fmt"
	"go-kafka-messaging/internal/app/receiver/client"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageReceiver struct {
	kafkaClient *client.KafkaClient
}

func NewMessageReceiver(client *client.KafkaClient) *MessageReceiver {
	return &MessageReceiver{kafkaClient: client}
}

func (self *MessageReceiver) StartReceive(topic string, channel chan<- kafka.Message) {

	consumer := self.kafkaClient.Get()

	if nil == consumer {
		panic("Client cannot be null")
	}

	err := consumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		fmt.Printf("Failed to subscribt topic: %s\n", err)
		os.Exit(1)
	}

	go func(consumer *kafka.Consumer, channel chan<- kafka.Message) {
		fmt.Println("Consumer Started")
		for {
			event, err := consumer.ReadMessage(10 * time.Second)

			if err != nil {
				continue
			}

			channel <- *event

		}
	}(consumer, channel)

}
