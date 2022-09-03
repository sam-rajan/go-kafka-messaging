package sender

type InputReader interface {
	ReadMessage()
}

type ReaderListener interface {
	OnInputRead(message string)
}

type InputHandler func(string)

func (f InputHandler) OnInputRead(message string) {
	f(message)
}
