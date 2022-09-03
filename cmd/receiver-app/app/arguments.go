package app

import (
	"log"
	"os"
	"strconv"
)

func GetConsumerCount() int {
	numberOfConsumers, _ := strconv.Atoi(os.Args[1])
	log.Printf("Consumer count received from CMD is %d\n", numberOfConsumers)
	if numberOfConsumers == 0 {
		numberOfConsumers = 2
	}

	return numberOfConsumers
}
