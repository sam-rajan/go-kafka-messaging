package main

import (
	"go-kafka-messaging/cmd/receiver-app/app"
	"sync"
)

func main() {
	consumerCount := app.GetConsumerCount()
	var wg = &sync.WaitGroup{}
	app.StartConsumers(consumerCount, wg)
	wg.Wait()
}
