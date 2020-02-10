package message

type Event struct {
	Type byte
	Name string
	Addr string
	Port uint16
	Msgs []*Message
}
