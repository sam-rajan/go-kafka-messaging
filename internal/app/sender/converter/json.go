package converter

import (
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"io/ioutil"
	"log"

	"github.com/riferrei/srclient"
)

type JsonDataFormat struct {
	registry *schemaregistry.SchemaRegistry
}

func NewJsonDataFormat(schemaRegistry *schemaregistry.SchemaRegistry) *JsonDataFormat {
	return &JsonDataFormat{registry: schemaRegistry}
}

func (self *JsonDataFormat) Convert(message []byte, schemaName string) ([]byte, int) {
	schema, err := self.registry.GetClient().GetLatestSchema(schemaName)
	if schema == nil {
		schemaBytes, _ := ioutil.ReadFile("assets/message.json")
		schema, err = self.registry.GetClient().CreateSchema(schemaName, string(schemaBytes), srclient.Json)
		if err != nil {
			log.Fatalf("Error creating the schema %s", err)
		}
	}

	return message, schema.ID()
}
