package main

import (
	"bufio"
	"fmt"
	inputparser "go-kafka-messaging/internal/app/sender/input-parser"
	"os"
)

func main() {

	// messageSender := app.Init()

	// topic := "first_topic"
	scanner := bufio.NewScanner(os.Stdin)
	var i int = 0
	for scanner.Scan() {

		value := scanner.Text()

		if value == "" {
			continue
		}

		//key := "key-" + strconv.Itoa(i)

		message, err := inputparser.ParseMessage(value)
		if err != nil {
			fmt.Printf("Input Parsing failed. Reason : %s\n", err.Error())
			continue
		}

		fmt.Println(message.GetMessage())
		fmt.Println(message.GetTopic())
		// message := kafka.Message{
		// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		// 	Key:            []byte(key),
		// 	Value:          []byte(value),
		// }
		// messageSender.Send(message)
		i = i + 1
	}
}
