package handler

import (
	"errors"
	"go-kafka-messaging/internal/app/sender"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	topicmanager "go-kafka-messaging/internal/pkg/topic-manager"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var messageSender *sender.KafkaMessageSender
var messageCounter = 0

func init() {
	log.Println("Initializing Sender Handler")
	configFile := "configs/kafka-producer.properties"
	configMap := configreader.ReadConfig(configFile)
	registryConfigFile := "configs/kafka-registry.properties"
	registryConfigMap := configreader.ReadConfig(registryConfigFile)

	if err := initializeBroadcastTopic(configMap); err != nil {
		log.Fatalln("Failed to initialize app")
	}
	registryClient := schemaregistry.CreateSchemaRegistry(registryConfigMap)
	messageSender = sender.NewMessageSender(configMap, registryClient)
}

func initializeBroadcastTopic(kafkaProperties kafka.ConfigMap) error {
	topicCreator, err := topicmanager.NewTopicClient[topicmanager.TopicCreator](kafkaProperties)

	if err != nil {
		log.Println("Failed To Create Topic Creator instance")
		return errors.New("Failed to create admin client")
	}

	_, err = topicCreator.CreateTopic("all")

	if err != nil {
		log.Printf("Faile to create Broadcast Topic , Reason: %s\n", err)
		return errors.New("Failed to create Broadcast Topic")
	}

	return nil
}
