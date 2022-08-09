package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
}

func (self *KafkaConsumer) Consume(config kafka.ConfigMap) {

	config["group.id"] = "first-consumer"
	config["auto.offset.reset"] = "earliest"

	topic := "first_topic"
	consumer, err := kafka.NewConsumer(&config)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	err = consumer.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		fmt.Printf("Failed to subscribt topic: %s", err)
		os.Exit(1)
	}

	for {
		event, err := consumer.ReadMessage(10 * time.Second)

		if err != nil {
			continue
		}

		fmt.Printf("Key = %s , Value = %s", string(event.Key), string(event.Value))
	}
}
