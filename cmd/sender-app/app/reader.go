package app

import (
	"bufio"
	"go-kafka-messaging/internal/app/sender"
	"os"
)

type ConsoleInputReader struct {
}

func NewInputReader() sender.InputReader {
	return &ConsoleInputReader{}
}

func (self *ConsoleInputReader) ReadMessage(listener sender.ReaderListener) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		value := scanner.Text()

		if value == "" {
			continue
		}

		if nil != listener {
			listener.OnInputRead(value)
		}

	}
}
