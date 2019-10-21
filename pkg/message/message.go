package message

import "encoding/json"

type Message struct {
	Version byte       `msgp:"v" json:"v"`
	Type    MsgControl `msgp:"t" json:"t"`
	Topic   string     `msgp:"tp" json:"tp"`
	ID      string     `msgp:"i" json:"i"`
	Payload []byte     `msgp:"p" json:"p"`
}

func New() *Message {
	return &Message{}
}

func (m *Message) Encode() ([]byte, error) {
	body, err := json.Marshal(m)
	return body, err
}

func (m *Message) Decode(body []byte) error {
	err := json.Unmarshal(body, m)
	return err
}
