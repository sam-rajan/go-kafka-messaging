package app

import (
	"go-kafka-messaging/internal/app/receiver"
	"go-kafka-messaging/internal/app/receiver/client"
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
	"log"
	"strconv"
	"sync"
)

func StartConsumers(numberOfConsumers int, wg *sync.WaitGroup) {

	schemaRegistryConfig, err := client.GetRegistryConfig()

	if nil != err {
		log.Printf("Failed to get registry config, Reason: %s \n", err)
		return
	}

	registryClient := schemaregistry.CreateSchemaRegistry(schemaRegistryConfig)

	for i := 0; i < numberOfConsumers; i++ {

		groupId := "Receiver-" + strconv.Itoa(i)
		kafkaTopic, err := client.CreateTopic(groupId)

		if nil != err {
			log.Printf("Failed to start consumer %s , Reason: %s \n", groupId, err)
			return
		}

		config, err := client.GetConfig()
		if nil != err {
			log.Printf("Failed to start consumer %s , Reason: %s \n", groupId, err)
			return
		}

		config.Set("group.id=" + groupId)
		messageReceiver := receiver.NewMessageReceiver(config)
		messageReceiver.StartReceive(kafkaTopic, wg)
		receiver.ProcessMessage(kafkaTopic, registryClient)
	}
}
