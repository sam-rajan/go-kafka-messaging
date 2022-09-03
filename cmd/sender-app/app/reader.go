package app

import (
	"bufio"
	"go-kafka-messaging/internal/app/sender"
	"os"
)

type ConsoleInput struct {
	listener       sender.ReaderListener
	currentMessage string
}

func NewConsoleInputReader(inputListener sender.ReaderListener) sender.InputReader {
	return &ConsoleInput{listener: inputListener}
}

func (self *ConsoleInput) ReadMessage() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		value := scanner.Text()

		if value == "" {
			continue
		}

		self.currentMessage = value
		if nil != self.listener {
			self.listener.OnInputRead(value)
		}

	}
}
