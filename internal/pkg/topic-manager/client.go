package topicmanager

import (
	"fmt"
	configreader "go-kafka-messaging/internal/pkg/config-reader"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type TopicClient struct {
	adminClient *kafka.AdminClient
}

func NewTopicClient(properties configreader.KafkaProperties) (*TopicClient, error) {
	delete(properties.Value, "auto.offset.reset")
	client, err := kafka.NewAdminClient(&properties.Value)

	if err != nil {
		return nil, err
	}

	return &TopicClient{adminClient: client}, nil
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
