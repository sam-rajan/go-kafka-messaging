package receiver

import (
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
}

func (self *KafkaProcessor) Process(wg *sync.WaitGroup, channel <-chan kafka.Message) {

	go func(wg *sync.WaitGroup, channel <-chan kafka.Message) {
		wg.Add(1)
		fmt.Println("Processor Started")
		for event := range channel {
			fmt.Printf("Received Key = %s , Value = %s \n ", string(event.Key), string(event.Value))
		}
	}(wg, channel)

}
