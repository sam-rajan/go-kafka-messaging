package app

import (
	"go-kafka-messaging/internal/app/receiver"
	"go-kafka-messaging/internal/app/receiver/client"
	"log"
	"strconv"
	"sync"
)

func StartConsumers(numberOfConsumers int, wg *sync.WaitGroup) {

	for i := 0; i < numberOfConsumers; i++ {

		groupId := "Receiver-" + strconv.Itoa(i)
		kafkaTopic, err := client.CreateTopic(groupId)

		if nil != err {
			log.Printf("Failed to start consumer %s , Reason: %s \n", groupId, err)
			return
		}

		config, err := client.GetConfig()
		if nil != err {
			log.Printf("Failed to start consumer %s , Reason: %s \n", groupId, err)
			return
		}

		config.Set("group.id=" + groupId)
		messageReceiver := receiver.NewMessageReceiver(config)
		messageReceiver.StartReceive(kafkaTopic, wg)
		receiver.ProcessMessage(kafkaTopic)
	}
}
