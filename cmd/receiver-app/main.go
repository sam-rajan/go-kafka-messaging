package main

import (
	"fmt"
	"go-kafka-messaging/cmd/receiver-app/app"
	receiver "go-kafka-messaging/internal/app/receiver"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
	topicmanager "go-kafka-messaging/internal/pkg/topic-manager"
	"os"
	"strconv"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	numberOfConsumers, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Consumer count received from CMD is %d\n", numberOfConsumers)
	if numberOfConsumers == 0 {
		numberOfConsumers = 2
	}

	var wg = &sync.WaitGroup{}
	wg.Add(1)

	configFile := "configs/kafka-consumer.properties"
	kafkaProperties := configreader.KafkaProperties{}
	kafkaProperties.LoadProperties(configFile)

	topicCreator, err := topicmanager.NewTopicClient[topicmanager.TopicCreator](kafkaProperties)

	if err != nil {
		fmt.Println("Failed To Create Topic Creator instance")
	}

	for i := 0; i < numberOfConsumers; i++ {
		channel := make(chan kafka.Message, 5)

		groupId := "Receiver-" + strconv.Itoa(i)
		_, err = topicCreator.CreateTopic(groupId)

		if err != nil {
			fmt.Printf("Faile to create Topic : %s , Reason: %s\n", groupId, err)
		}

		kafkaProperties.Value.Set("group.id=" + groupId)

		messageReceiver := app.Init(kafkaProperties)
		processor := receiver.KafkaProcessor{}
		processor.Process(channel)
		messageReceiver.StartReceive(groupId, channel)
	}

	wg.Wait()
}
