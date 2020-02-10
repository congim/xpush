package msgid

import "go.uber.org/zap"

type MsgID interface {
	MsgID() (string, error)
}

func New(l *zap.Logger) MsgID {
	return newNoopMsgID()
}
