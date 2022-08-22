package app

import (
	"fmt"
	"go-kafka-messaging/internal/app/sender"
	"go-kafka-messaging/internal/app/sender/client"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	topicmanager "go-kafka-messaging/internal/pkg/topic-manager"
)

func Init() *sender.MessageSender {
	configFile := "configs/kafka-producer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	topicCreator, err := topicmanager.NewTopicClient[topicmanager.TopicCreator](kafkaProperties)

	if err != nil {
		fmt.Println("Failed To Create Topic Creator instance")
	}

	_, err = topicCreator.CreateTopic("all")

	if err != nil {
		fmt.Printf("Faile to create Broadcast Topic , Reason: %s\n", err)
	}

	kafkaClient := client.KafkaClient{}
	kafkaClient.Init(&kafkaProperties)

	messageSender := sender.NewMessageSender(&kafkaClient)
	return messageSender
}
