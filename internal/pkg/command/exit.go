package command

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ExitCommand struct {
	data string
}

func NewExitCommand(message string) Command {
	return &TextCommand{data: message}
}

func (self *ExitCommand) Execute() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		log.Println("Sutting down gracefully.")
		// clean up here
		os.Exit(0)
	}()
}
