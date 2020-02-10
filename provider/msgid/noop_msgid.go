package msgid

import (
	"fmt"
	"time"
)

var _ MsgID = (*noopMsgID)(nil)

type noopMsgID struct {
}

func newNoopMsgID() *noopMsgID {
	return &noopMsgID{}
}

func (noopMsgID *noopMsgID) MsgID() (string, error) {
	return fmt.Sprintf("%d", time.Now().UnixNano()), nil
}
