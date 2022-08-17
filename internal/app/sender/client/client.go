package client

import (
	"fmt"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaClient struct {
	instance *kafka.Producer
}

func (self *KafkaClient) Get() *kafka.Producer {
	return self.instance
}

func (self *KafkaClient) Init(config *configreader.KafkaProperties) {

	producer, err := kafka.NewProducer(&config.Value)

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	go func(producer *kafka.Producer) {
		fmt.Println("Producer Callback started")
		for event := range producer.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}(producer)

	self.instance = producer
}
