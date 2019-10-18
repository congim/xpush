package message

type MsgControl byte

const (
	_ MsgControl = iota
	MsgPub
	MsgPull
)

//type MsgVersion byte
//
//const (
//	_ MsgVersion = iota
//)
