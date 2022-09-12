package sender

import (
	"encoding/binary"
	"go-kafka-messaging/internal/app/sender/converter"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const SCHEMA string = "MessageSchema"

type KafkaMessageSender struct {
	instance       *kafka.Producer
	schemaRegistry *schemaregistry.SchemaRegistry
	messageFormat  string
}

func NewMessageSender(config kafka.ConfigMap,
	registry *schemaregistry.SchemaRegistry, dataFormat string) *KafkaMessageSender {
	producer, err := kafka.NewProducer(&config)

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	sender := &KafkaMessageSender{
		instance:       producer,
		schemaRegistry: registry,
		messageFormat:  dataFormat,
	}
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

		defer producer.Close()
	}(self.instance)
}

func (self *KafkaMessageSender) Send(reciever string, key string, message []byte) {
	log.Println("Sending Message")

	if nil == self.instance {
		panic("No Kafka producer instance found")
	}

	dataConverter := converter.GetDataConverter(self.messageFormat, self.schemaRegistry)
	schemaName := SCHEMA + "_" + self.messageFormat
	convertedMessage, schemaId := dataConverter.Convert(message, schemaName)

	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schemaId))
	schemaIdHeader := kafka.Header{Key: "schemaId", Value: schemaIDBytes}
	dataFormatHeader := kafka.Header{Key: "dataFormat", Value: []byte(self.messageFormat)}

	kafkaMessage := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &reciever, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          convertedMessage,
		Headers:        []kafka.Header{schemaIdHeader, dataFormatHeader},
	}
	self.instance.Produce(&kafkaMessage, nil)
}
