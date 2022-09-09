package sender

import (
	"encoding/binary"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"io/ioutil"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/riferrei/srclient"
)

const SCHEMA string = "MessageSchema"

type KafkaMessageSender struct {
	instance       *kafka.Producer
	schemaRegistry *schemaregistry.SchemaRegistry
}

func NewMessageSender(config kafka.ConfigMap, registry *schemaregistry.SchemaRegistry) *KafkaMessageSender {
	producer, err := kafka.NewProducer(&config)

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	sender := &KafkaMessageSender{instance: producer, schemaRegistry: registry}
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

	schema, err := self.schemaRegistry.GetClient().GetLatestSchema(SCHEMA)
	if schema == nil {
		schemaBytes, _ := ioutil.ReadFile("assets/message.json")
		schema, err = self.schemaRegistry.GetClient().CreateSchema(SCHEMA, string(schemaBytes), srclient.Json)
		if err != nil {
			log.Fatalf("Error creating the schema %s", err)
		}
	}

	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.ID()))
	header := kafka.Header{Key: "schemaId", Value: schemaIDBytes}

	kafkaMessage := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &reciever, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          message,
		Headers:        []kafka.Header{header},
	}
	self.instance.Produce(&kafkaMessage, nil)
}
