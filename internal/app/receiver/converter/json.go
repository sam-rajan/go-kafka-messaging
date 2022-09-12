package converter

import (
	"encoding/json"
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"log"
)

type JsonDataFormat struct {
	registry *schemaregistry.SchemaRegistry
}

func NewJsonDataFormat(schemaRegistry *schemaregistry.SchemaRegistry) *JsonDataFormat {
	return &JsonDataFormat{registry: schemaRegistry}
}

func (self *JsonDataFormat) Convert(message []byte, schemaId uint32) (*inputparser.Message, error) {

	schema, err := self.registry.GetClient().GetSchema(int(schemaId))
	if err != nil {
		log.Fatalf("Error getting the schema with id '%d' %s", schemaId, err)
	}

	var jsonType interface{}
	if err := json.Unmarshal(message, &jsonType); err != nil {
		log.Printf("Failed to unmarshall Error: %v \n", err)
		return nil, err
	}

	jsonSchema := schema.JsonSchema()
	if err = jsonSchema.Validate(jsonType); err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	if nil != err {
		log.Printf("Failed unmarshling message. Reason : %s", err)
		return nil, err
	}

	parsedObject := &inputparser.Message{}
	if err := json.Unmarshal(message, parsedObject); err != nil {
		log.Printf("Failed to unmarshall to Message struct Error: %v \n", err)
		return nil, err
	}

	return parsedObject, nil
}
