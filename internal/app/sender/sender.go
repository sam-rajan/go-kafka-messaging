package sender

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaMessageSender struct {
	instance *kafka.Producer
}

func NewMessageSender(config kafka.ConfigMap) *KafkaMessageSender {
	producer, err := kafka.NewProducer(&config)

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	sender := &KafkaMessageSender{instance: producer}
	sender.initProducer()
	return sender
}

func (self *KafkaMessageSender) initProducer() {
	if nil == self.instance {
		panic("No Kafka producer instance found")
	}

	go func(producer *kafka.Producer) {
		log.Println("Producer Callback started")
		for event := range producer.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: Consumer Not Exist %s\n", *ev.TopicPartition.Topic)
				} else {
					log.Printf("Produced event to topic %s: key = %s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}(self.instance)
}

func (self *KafkaMessageSender) Send(message kafka.Message) {
	log.Println("Sending Message")

	if nil == self.instance {
		panic("No Kafka producer instance found")
	}

	self.instance.Produce(&message, nil)
}
