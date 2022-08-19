package topicmanager

type TopicFetcher interface {
	FetchTopics() []string
}

type TopicCreator interface {
	CreateTopic(topicName string) string
}
