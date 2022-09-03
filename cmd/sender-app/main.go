package main

import (
	"go-kafka-messaging/cmd/sender-app/app"
	"go-kafka-messaging/internal/app/sender"
	"go-kafka-messaging/internal/app/sender/handler"
)

func main() {
	//Listener which notifies message handler on every input
	listener := sender.InputHandler(handler.OnInputRead)
	//create console input reader and starts reading message
	inputReader := app.NewConsoleInputReader(listener)
	inputReader.ReadMessage()
}
