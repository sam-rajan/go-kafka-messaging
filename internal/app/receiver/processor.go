package receiver

import (
	"encoding/binary"
	"encoding/json"
	"go-kafka-messaging/internal/app/receiver/client"
	"go-kafka-messaging/internal/app/receiver/command"
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProcessMessage(topic *client.KafkaTopic, schemaRegistry *schemaregistry.SchemaRegistry) {

	go func(channel <-chan kafka.Message, schemaRegistry *schemaregistry.SchemaRegistry) {
		for event := range channel {
			topic.MessageCount = topic.MessageCount + 1

			if len(event.Headers) == 0 {
				log.Println("Invalid message. Schema ID header missing")
				continue
			}
			schemaId := binary.BigEndian.Uint32(event.Headers[0].Value)
			schema, err := schemaRegistry.GetClient().GetSchema(int(schemaId))
			if err != nil {
				log.Fatalf("Error getting the schema with id '%d' %s", schemaId, err)
			}

			var jsonType interface{}
			if err := json.Unmarshal(event.Value, &jsonType); err != nil {
				log.Printf("Failed to unmarshall Error: %v \n", err)
				continue
			}

			jsonSchema := schema.JsonSchema()
			if err = jsonSchema.Validate(jsonType); err != nil {
				log.Printf("%v", err)
				continue
			}

			if nil != err {
				log.Printf("Failed unmarshling message. Reason : %s", err)
				return
			}

			message := &inputparser.Message{}
			if err := json.Unmarshal(event.Value, message); err != nil {
				log.Printf("Failed to unmarshall to Message struct Error: %v \n", err)
				continue
			}
			message.Receiver = topic.Name
			commandFn := command.GetCommand(message.Type)
			if message.Type == command.EXIT {
				commandFn(topic)
			} else {
				commandFn(*message)
			}

		}
	}(topic.Channel, schemaRegistry)

}
