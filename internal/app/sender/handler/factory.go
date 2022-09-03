package handler

import (
	"go-kafka-messaging/internal/pkg/shutdown"
)

const (
	TEXT    = "TEXT"
	HISTORY = "HISTORY"
	EXIT    = "EXIT"
)

func GetCommand(action string) func(input interface{}) {

	switch action {
	default:
		return sendMessage
	case HISTORY:
		return getHistory
	case EXIT:
		return shutdown.GracefulShutdown
	}
}
