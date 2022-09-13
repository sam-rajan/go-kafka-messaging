package app

import (
	"go-kafka-messaging/internal/app/sender/converter"
	"log"
	"os"
	"strings"
)

func GetDataFormat() string {

	if len(os.Args) == 1 {
		log.Fatalln("No arguments supplied. Terminating application")
	}
	dataFormat := strings.ToUpper(os.Args[1])

	if dataFormat != converter.JSON &&
		dataFormat != converter.AVRO &&
		dataFormat != converter.PROTOBUF {
		log.Fatalf("Invalid data format %s \n", dataFormat)
	}
	log.Printf("Dataformat opted : %s\n", dataFormat)
	return dataFormat
}
