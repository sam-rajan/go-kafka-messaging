package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
}

func (self *KafkaProducer) Produce(config kafka.ConfigMap, wg *sync.WaitGroup) {

	topic := "first_topic"
	producer, err := kafka.NewProducer(&config)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	go func(producer *kafka.Producer, wg *sync.WaitGroup) {
		fmt.Println("Producer Callback started")
		wg.Add(1)
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
		wg.Done()
	}(producer, wg)

	go func(producer *kafka.Producer, wg *sync.WaitGroup) {
		fmt.Println("Producer Started")
		wg.Add(1)
		for {
			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Key:            []byte("nic-key"),
				Value:          []byte("New Data" + time.Now().String()),
			}, nil)

			time.Sleep(10 * time.Second)
		}
		wg.Done()
	}(producer, wg)

}
