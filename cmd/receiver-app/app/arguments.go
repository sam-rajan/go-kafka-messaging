package app

import (
	"log"
	"os"
	"strconv"
)

func GetConsumerCount() int {
	if len(os.Args) < 1 {
		log.Fatalln("No arguments supplied. Terminating application")
	}
	numberOfConsumers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid argument supplied for consumer count. Error : %v", err)
	}
	log.Printf("Consumer count received from CMD is %d\n", numberOfConsumers)
	if numberOfConsumers == 0 {
		log.Println("Falling back to 2 consumers since supplied consumer count is 0")
		numberOfConsumers = 2
	}

	return numberOfConsumers
}
