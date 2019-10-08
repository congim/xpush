package message

type Message struct {
	Version byte
	Type    byte
	ID      string
	Payload []byte
}
