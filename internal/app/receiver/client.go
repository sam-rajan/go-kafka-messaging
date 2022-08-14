package receiver

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Instance *kafka.Consumer
}

func (self *KafkaConsumer) Init(config kafka.ConfigMap) {
	consumer, err := kafka.NewConsumer(&config)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	self.Instance = consumer
}
