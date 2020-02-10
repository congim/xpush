package msgid

var _ MsgID = (*noopMsgID)(nil)

type noopMsgID struct {
}

func newNoopMsgID() *noopMsgID {
	return &noopMsgID{}
}

func (noopMsgID *noopMsgID) MsgID() (string, error) {
	return "", nil
}
