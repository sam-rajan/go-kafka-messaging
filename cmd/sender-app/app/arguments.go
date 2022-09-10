package app

import (
	"log"
	"os"
)

func getDataFormat() string {

	if len(os.Args) < 1 {
		log.Fatalln("No arguments supplied. Terminating application")
	}
	dataFormat := os.Args[1]
	log.Printf("Dataformat opted : %s\n", dataFormat)
	return dataFormat
}
