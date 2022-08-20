package topicmanager

import (
	configreader "go-kafka-messaging/internal/pkg/config-reader"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type TopicFetcher interface {
	FetchTopics() []string
}

type TopicCreator interface {
	CreateTopic(topicName string) (string, error)
}

func NewTopicClient[T any](properties configreader.KafkaProperties) (T, error) {
	delete(properties.Value, "auto.offset.reset")
	client, err := kafka.NewAdminClient(&properties.Value)

	if err != nil {
		var empty T
		return empty, err
	}

	return any(&TopicClient{adminClient: client}).(T), nil
}
