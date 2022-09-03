package topicmanager

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type TopicFetcher interface {
	FetchTopics() []string
}

type TopicCreator interface {
	CreateTopic(topicName string) (string, error)
}

func NewTopicClient[T any](properties kafka.ConfigMap) (T, error) {
	delete(properties, "auto.offset.reset")
	delete(properties, "group.id")
	client, err := kafka.NewAdminClient(&properties)

	if err != nil {
		var empty T
		return empty, err
	}

	return any(&TopicClient{adminClient: client}).(T), nil
}
