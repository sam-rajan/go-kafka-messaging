package receiver

import (
	"encoding/json"
	"go-kafka-messaging/internal/app/receiver/client"
	"go-kafka-messaging/internal/app/receiver/command"
	inputparser "go-kafka-messaging/internal/pkg/input-parser"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProcessMessage(topic *client.KafkaTopic) {

	go func(channel <-chan kafka.Message) {
		for event := range channel {
			topic.MessageCount = topic.MessageCount + 1
			message := &inputparser.Message{}
			err := json.Unmarshal(event.Value, message)

			if nil != err {
				log.Printf("Failed unmarshling message. Reason : %s", err)
				return
			}

			commandFn := command.GetCommand(message.Type)
			commandFn(*message)
		}
	}(topic.Channel)

}
