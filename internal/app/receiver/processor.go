package receiver

import (
	"encoding/binary"
	"go-kafka-messaging/internal/app/receiver/client"
	"go-kafka-messaging/internal/app/receiver/command"
	"go-kafka-messaging/internal/app/receiver/converter"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProcessMessage(topic *client.KafkaTopic, schemaRegistry *schemaregistry.SchemaRegistry) {

	go func(channel <-chan kafka.Message, schemaRegistry *schemaregistry.SchemaRegistry) {
		for event := range channel {
			topic.MessageCount = topic.MessageCount + 1

			schemaId, dataFormat := uint32(0), ""
			for _, header := range event.Headers {
				switch header.Key {
				case "schemaId":
					schemaId = binary.BigEndian.Uint32(header.Value)
				case "dataFormat":
					dataFormat = string(header.Value)
				}
			}

			if schemaId == 0 || dataFormat == "" {
				log.Printf("Missing headers, cannot parse message")
				continue
			}

			dataConverter := converter.GetDataConverter(dataFormat, schemaRegistry)
			message, err := dataConverter.Convert(event.Value, schemaId)

			if err != nil {
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
