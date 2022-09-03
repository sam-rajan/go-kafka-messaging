package shutdown

import (
	"log"
	"os"
)

func GracefulShutdown(interface{}) {
	go func() {
		log.Println("Sutting down gracefully.")
		// clean up here
		os.Exit(0)

	}()
}
