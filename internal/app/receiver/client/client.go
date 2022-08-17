package client

import (
	"fmt"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaClient struct {
	instance *kafka.Consumer
}

func (self *KafkaClient) Get() *kafka.Consumer {
	return self.instance
}

func (self *KafkaClient) Init(config *configreader.KafkaProperties) {
	consumer, err := kafka.NewConsumer(&config.Value)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	self.instance = consumer
}
