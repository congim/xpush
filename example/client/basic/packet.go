package basic

//PushPacket pusMsg struct
type PushPacket struct {
	TTopics   []string `msg:"tts"`   // to topic
	TOpIDs    []string `msg:"tpids"` // to user
	FromOpID  string   `msg:"fid"`   // from operatorid
	FromTopic string   `msg:"ftop"`  // from topic
	PackType  int      `msg:"pt"`    // packtype
	MsgID     string   `msg:"mid"`   // msgid
	IosMsg    string   `msg:"i"`     //
	Expire    int64    `msg:"e"`     // time out
	Label     string   `msg:"l"`     // label
	MsgType   int      `msg:"mt"`    // msgtype
	Message   []byte   `msg:"ms"`    // 消息体
	IosJSON   []byte   `msg:"ij"`    // iosjson
	Nick      string   `msg:"nk"`    // Nick
}

// NewPushPacket new push packet
type NewPushPacket struct {
	Badge string      `msg:"badge"`
	Data  *PushPacket `msg:"data"`
}

// MsgType   int32    `msg:"mt"`    // 群聊内部消息类型

// PushMessage 消息推送结构
type PushMessage struct {
	MsgType   int      `msg:"mt"`
	SessionID string   `msg:"id"`
	Topic     string   `msg:"tc"`
	Message   [][]byte `msg:"m"`
}
