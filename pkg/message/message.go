package message

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/congim/xpush/pkg/tool"
	"github.com/golang/snappy"
)

// Message msg
type Message struct {
	Type    byte   `msgp:"t" json:"t"`   // 消息类型
	Topic   string `msgp:"tp" json:"tp"` // 主题
	ID      string `msgp:"i" json:"i"`   // 消息ID
	Payload []byte `msgp:"p" json:"p"`   // 具体消息体
}

// Encode encode
func (m *Message) Encode() ([]byte, error) {
	body, err := json.Marshal(m)
	return body, err
}

// Decode decode
func (m *Message) Decode(body []byte) error {
	err := json.Unmarshal(body, m)
	return err
}

// New new
func New() *Message {
	return &Message{}
}

// Encode encode
func Encode(msgs []*Message, isCompress byte) ([]byte, error) {
	body, err := json.Marshal(msgs)
	if err != nil {
		return nil, err
	}
	var newBody []byte
	var compress byte = 0
	if isCompress == Compress {
		body = snappy.Encode(nil, body)
		compress, err = tool.SetBitValue(compress, 0, int(Compress))
	} else {
		compress, err = tool.SetBitValue(compress, 0, int(NoCompress))
	}

	if err != nil {
		return nil, err
	}

	newBody = make([]byte, len(body)+1)
	newBody[0] = compress
	copy(newBody[1:], body)
	return newBody, nil
}

// Decode decode
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

func PackPullMsg(offset int, checkTime int64) []byte {
	msg := make([]byte, 12)
	binary.PutUvarint(msg[0:3], uint64(offset))
	binary.PutUvarint(msg[3:11], uint64(checkTime))
	//binary.PutUvarint(msg[3:11], uint64(checkTime))
	return msg
}

func UnPackPullMsg(b []byte) (int, int64, error) {
	if len(b) != 12 {
		return 0, 0, fmt.Errorf("报文长度错误")
	}
	offset, _ := binary.Uvarint(b[0:3])
	checkTime, _ := binary.Uvarint(b[3:11])
	return int(offset), int64(checkTime), nil
}
