package converter

import (
	"encoding/json"
	"fmt"
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"log"
)

type AvroDataFormat struct {
	registry *schemaregistry.SchemaRegistry
}

func NewAvroDataFormat(schemaRegistry *schemaregistry.SchemaRegistry) *AvroDataFormat {
	return &AvroDataFormat{registry: schemaRegistry}
}

func (self *AvroDataFormat) Convert(message []byte, schemaId uint32) (*inputparser.Message, error) {
	schema, err := self.registry.GetClient().GetSchema(int(schemaId))
	if err != nil {
		log.Printf("Error getting the schema with id '%d' %s", schemaId, err)
		return nil, err
	}
	fmt.Println("Hello")
	native, _, _ := schema.Codec().NativeFromBinary(message)
	value, _ := schema.Codec().TextualFromNative(nil, native)

	parsedObject := &inputparser.Message{}
	if err := json.Unmarshal(value, parsedObject); err != nil {
		log.Printf("Failed to unmarshall to Message struct Error: %v \n", err)
		return nil, err
	}

	return parsedObject, nil
}
