package app

import (
	"errors"
	"fmt"
	"go-kafka-messaging/internal/app/sender"
	"go-kafka-messaging/internal/app/sender/client"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	topicmanager "go-kafka-messaging/internal/pkg/topic-manager"
)

func InitializeSender(kafkaProperties *configreader.KafkaProperties) *sender.MessageSender {

	kafkaClient := client.KafkaClient{}
	kafkaClient.Init(kafkaProperties)

	messageSender := sender.NewMessageSender(&kafkaClient)
	return messageSender
}

func InitializeBroadcastTopic(kafkaProperties configreader.KafkaProperties) error {
	topicCreator, err := topicmanager.NewTopicClient[topicmanager.TopicCreator](kafkaProperties)

	if err != nil {
		fmt.Println("Failed To Create Topic Creator instance")
		return errors.New("Failed to create admin client")
	}

	_, err = topicCreator.CreateTopic("all")

	if err != nil {
		fmt.Printf("Faile to create Broadcast Topic , Reason: %s\n", err)
		return errors.New("Failed to create Broadcast Topic")
	}

	return nil
}
