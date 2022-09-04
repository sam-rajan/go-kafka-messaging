package handler

import (
	"log"
	"os"
)

func shutdown(interface{}) {
	go func() {
		log.Println("Sutting down application.")
		// clean up here
		os.Exit(0)

	}()
}
