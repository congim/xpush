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

const (
	MAX_MESSAGE_PULL_COUNT int = 200
)

var (
	MSG_NEWEST_OFFSET = []byte("0")
)
