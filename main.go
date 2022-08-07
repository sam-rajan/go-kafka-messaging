package main

import (
	"fmt"
)

func main() {
	kafkaProperties := KafkaProperties{}
	kafkaProperties.LoadProperties()

	fmt.Printf(kafkaProperties.bootstrapServer)
}
