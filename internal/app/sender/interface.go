package sender

import "github.com/confluentinc/confluent-kafka-go/kafka"

type MessageSender interface {
	Send(message kafka.Message)
}

type InputReader interface {
	ReadMessage(listener ReaderListener)
}

type ReaderListener interface {
	OnInputRead(message string)
}
