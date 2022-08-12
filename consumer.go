package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
}

func (self *KafkaConsumer) Consume(config kafka.ConfigMap, wg *sync.WaitGroup, channel chan<- kafka.Message) {

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

	go func(consumer *kafka.Consumer, wg *sync.WaitGroup) {
		wg.Add(1)
		fmt.Println("Consumer Started")
		for {
			event, err := consumer.ReadMessage(10 * time.Second)

			if err != nil {
				continue
			}

			channel <- *event

		}
		wg.Done()
	}(consumer, wg)

}
