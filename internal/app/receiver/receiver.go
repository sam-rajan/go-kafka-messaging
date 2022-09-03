package receiver

import (
	"go-kafka-messaging/internal/app/receiver/client"
	"log"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageReceiver struct {
	instance *kafka.Consumer
}

func NewMessageReceiver(config kafka.ConfigMap) *MessageReceiver {
	consumer, err := kafka.NewConsumer(&config)

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	return &MessageReceiver{instance: consumer}
}

func (self *MessageReceiver) StartReceive(topic *client.KafkaTopic, wg *sync.WaitGroup) {

	if nil == self.instance {
		panic("Client cannot be null")
	}

	topic.Channel = make(chan kafka.Message, 5)

	broadcastTopic := "all"
	err := self.instance.SubscribeTopics([]string{topic.Name, broadcastTopic}, nil)

	if err != nil {
		log.Fatalf("Failed to subscribt topic: %s\n", err)
	}

	go func(consumer *kafka.Consumer, channel chan<- kafka.Message) {
		log.Println("Consumer Started. Name: " + topic.Name)
		wg.Add(1)
		for {
			event, err := consumer.ReadMessage(10 * time.Second)

			if err != nil {
				continue
			}

			channel <- *event

		}
	}(self.instance, topic.Channel)

}
