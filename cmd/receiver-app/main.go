package main

import (
	"fmt"
	"go-kafka-messaging/cmd/receiver-app/app"
	receiver "go-kafka-messaging/internal/app/receiver"
	configreader "go-kafka-messaging/internal/pkg/config-reader"
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

	for i := 0; i < numberOfConsumers; i++ {
		channel := make(chan kafka.Message, 5)

		groupId := "Group-" + strconv.Itoa(i)
		kafkaProperties.Value.Set("group.id=" + groupId)

		messageReceiver := app.Init(kafkaProperties)
		topic := "first_topic"
		processor := receiver.KafkaProcessor{}
		processor.Process(channel)
		messageReceiver.StartReceive(topic, channel)
	}

	wg.Wait()
}
