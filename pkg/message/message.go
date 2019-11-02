package message

import (
	"encoding/json"

	"github.com/congim/xpush/pkg/tool"
	"github.com/golang/snappy"
)

type Message struct {
	Type    int    `msgp:"t" json:"t"`
	Topic   string `msgp:"tp" json:"tp"`
	ID      string `msgp:"i" json:"i"`
	Payload []byte `msgp:"p" json:"p"`
	From    string `-`
}

func (m *Message) Encode() ([]byte, error) {
	body, err := json.Marshal(m)
	return body, err
}

func (m *Message) Decode(body []byte) error {
	err := json.Unmarshal(body, m)
	return err
}

func New() *Message {
	return &Message{}
}

func Decode(body []byte) ([]*Message, error) {
	isCompress, err := tool.GetBitValue(body[0], 0)
	if err != nil {
		return nil, err
	}
	if isCompress == Compress {
		body, err = snappy.Decode(nil, body[1:])
		if err != nil {
			return nil, err
		}
	}

	var msgs []*Message
	if isCompress == Compress {
		err = json.Unmarshal(body, &msgs)
	} else {
		err = json.Unmarshal(body[1:], &msgs)
	}
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func Encode(msgs []*Message, isCompress byte) ([]byte, error) {
	body, err := json.Marshal(msgs)
	if err != nil {
		return nil, err
	}
	var newBody []byte
	var compress byte = 0
	if isCompress == Compress {
		body = snappy.Encode(nil, body)
		compress, err = tool.SetBitValue(compress, 0, 1)
	} else {
		compress, err = tool.SetBitValue(compress, 0, 0)
	}

	if err != nil {
		return nil, err
	}

	newBody = make([]byte, len(body)+1)
	newBody[0] = compress
	copy(newBody[1:], body)

	return newBody, nil
}
