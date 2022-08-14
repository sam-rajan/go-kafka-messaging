package receiver

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageReceiver struct {
	KafkaConsumer *kafka.Consumer
}

func (self *MessageReceiver) StartReceive(topic string, wg *sync.WaitGroup, channel chan<- kafka.Message) {

	err := self.KafkaConsumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		fmt.Printf("Failed to subscribt topic: %s\n", err)
		os.Exit(1)
	}

	go func(consumer *kafka.Consumer, wg *sync.WaitGroup, channel chan<- kafka.Message) {
		wg.Add(1)
		fmt.Println("Consumer Started")
		for {
			event, err := consumer.ReadMessage(10 * time.Second)

			if err != nil {
				continue
			}

			channel <- *event

		}
	}(self.KafkaConsumer, wg, channel)

}
