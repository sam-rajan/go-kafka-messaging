package receiver

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
}

func (self *KafkaProcessor) Process(channel <-chan kafka.Message) {

	go func(channel <-chan kafka.Message) {
		for event := range channel {
			fmt.Printf("Received message in %s consumer\n", *event.TopicPartition.Topic)
			fmt.Printf("Received Key = %s, Value = %s\n ", string(event.Key), string(event.Value))
		}
	}(channel)

}
