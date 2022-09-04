package client

import (
	topicmanager "go-kafka-messaging/internal/pkg/topic-manager"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaTopic struct {
	Name         string
	MessageCount int
	Channel      chan kafka.Message
	IsRunning    bool
}

func CreateTopic(groupId string) (*KafkaTopic, error) {

	topicCreator, err := topicmanager.NewTopicClient[topicmanager.TopicCreator](config)

	if err != nil {
		log.Println("Failed To Create Topic Creator instance")
		return nil, err
	}

	_, err = topicCreator.CreateTopic(groupId)

	if err != nil {
		log.Printf("Faile to create Topic : %s , Reason: %s\n", groupId, err)
		return nil, err
	}

	return &KafkaTopic{Name: groupId, IsRunning: true}, nil
}
