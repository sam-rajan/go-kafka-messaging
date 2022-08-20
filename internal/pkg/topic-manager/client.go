package topicmanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type TopicClient struct {
	adminClient *kafka.AdminClient
}

func (self TopicClient) FetchTopics() []string {
	metadata, err := self.adminClient.GetMetadata(nil, true, 10000)

	if err != nil {
		fmt.Println("Failed to fetch topic list")
		return nil
	}

	topics := make([]string, 0, len(metadata.Topics))
	for _, tm := range metadata.Topics {
		topics = append(topics, tm.Topic)
	}

	return topics

}

func (self TopicClient) CreateTopic(topicName string) (string, error) {
	existingTopics := self.FetchTopics()

	for _, topic := range existingTopics {
		if topic == topicName {
			return topic, nil
		}
	}

	topicSpec := kafka.TopicSpecification{Topic: topicName, NumPartitions: 2}
	topicSpecArr := []kafka.TopicSpecification{topicSpec}

	topicResult, err := self.adminClient.CreateTopics(context.TODO(), topicSpecArr)

	if err != nil {
		return "", err
	}

	if topicResult[0].Error.Code() != kafka.ErrNoError {
		return "", errors.New(topicResult[0].Error.Code().String())
	}

	return topicResult[0].Topic, nil
}
