package sender

import "github.com/confluentinc/confluent-kafka-go/kafka"

type MessageSender interface {
	Send(message kafka.Message)
}

type MessageReader interface {
	ReadMessage(listener ReaderListener)
}

type ReaderListener interface {
	onMessageRead(message string)
}
