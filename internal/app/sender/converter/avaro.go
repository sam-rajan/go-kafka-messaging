package converter

import (
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"io/ioutil"
	"log"

	"github.com/riferrei/srclient"
)

type AvaroDataFormat struct {
	registry *schemaregistry.SchemaRegistry
}

func NewAvaroDataFormat(schemaRegistry *schemaregistry.SchemaRegistry) *JsonDataFormat {
	return &JsonDataFormat{registry: schemaRegistry}
}

func (self *AvaroDataFormat) Convert(message []byte, schemaName string) ([]byte, int) {
	schema, err := self.registry.GetClient().GetLatestSchema(schemaName)
	if schema == nil {
		schemaBytes, _ := ioutil.ReadFile("assets/message.avsc")
		schema, err = self.registry.GetClient().CreateSchema(schemaName, string(schemaBytes), srclient.Avro)
		if err != nil {
			log.Fatalf("Error creating the schema %s", err)
		}
	}

	native, _, _ := schema.Codec().NativeFromTextual(message)
	valueBytes, _ := schema.Codec().BinaryFromNative(nil, native)

	return valueBytes, schema.ID()
}
