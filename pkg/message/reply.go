package message

type Reply struct {
	Message string `msgp:"m" json:"p"`
}

func NewReply() *Reply {
	return &Reply{}
}
