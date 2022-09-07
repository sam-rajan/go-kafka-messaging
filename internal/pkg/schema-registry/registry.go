package schemaregistry

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/riferrei/srclient"
)

type SchemaRegistry struct {
	client *srclient.SchemaRegistryClient
}

func CreateSchemaRegistry(config kafka.ConfigMap) *SchemaRegistry {

	registryUrl, err := config.Get("registry.url", "")

	if err != nil {
		log.Fatal("Failed to create schema registry client")
	}

	username, err := config.Get("registry.username", "")

	if err != nil {
		log.Fatal("Failed to create schema registry client")
	}

	password, err := config.Get("registry.pasword", "")

	if err != nil {
		log.Fatal("Failed to create schema registry client")
	}

	schemaRegistryClient := srclient.CreateSchemaRegistryClient(registryUrl.(string))
	schemaRegistryClient.SetCredentials(username.(string), password.(string))

	return &SchemaRegistry{client: schemaRegistryClient}
}
