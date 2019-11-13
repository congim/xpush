package message

type MsgControl byte

const (
	MsgPub  = 10000
	MsgPull = 10001
	Sub     = 10002
	UnSub   = 10003
	NewMsg  = 10004
)

//type MsgVersion byte
//
const (
	NoCompress byte = iota
	Compress
)
