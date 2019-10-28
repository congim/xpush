package message

type MsgControl byte

const (
	//_       int = iota
	MsgPub  = 10000
	MsgPull = 10001
	Sub     = 10002
	UnSub   = 10003
	//Login
	//Logout
)

//type MsgVersion byte
//
//const (
//	_ MsgVersion = iota
//)
