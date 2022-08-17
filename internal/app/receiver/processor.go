package receiver

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
}

func (self *KafkaProcessor) Process(channel <-chan kafka.Message) {

	go func(hannel <-chan kafka.Message) {
		for event := range channel {
			fmt.Printf("Received Key = %s , Value = %s \n ", string(event.Key), string(event.Value))
		}
	}(channel)

}
