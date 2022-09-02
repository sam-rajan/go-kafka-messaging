package app

import (
	"errors"
	"fmt"
	topicmanager "go-kafka-messaging/internal/pkg/topic-manager"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitializeBroadcastTopic(kafkaProperties kafka.ConfigMap) error {
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
