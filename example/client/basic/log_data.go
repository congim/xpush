package basic

// HTTPPushLog http
//easyjson:json
type HTTPPushLog struct {
	Topics  string `json:"topics"`
	MsgID   string `json:"msgID"`
	Iosmsg  string `json:"iosmsg"`
	Expire  string `json:"expire"`
	MsgType string `json:"msgType"`
	Message string `json:"message"`
	Label   string `json:"label"`
	From    string `json:"from"`
}

// InsertLog  insert redis data log
//easyjson:json
type InsertLog struct {
	Topic   string `json:"topic"`
	MsgID   string `json:"msgID"`
	MsgType uint8  `json:"msgType"`
	Label   string `json:"label"`
}

// CachePushLog cache push message to nats
//easyjson:json
type CachePushLog struct {
	Topic   string `json:"topic"`
	MsgID   string `json:"msgID"`
	MsgType uint8  `json:"msgType"`
	Label   string `json:"label"`
}

// RoomPushClientLog room push data to client
//easyjson:json
type RoomPushClientLog struct {
	Topic         string `json:"topic"`
	OriginalMsgID string `json:"originalMsgID"`
	NewMsgID      uint16 `json:"newMsgID"`
	MsgType       uint8  `json:"msgType"`
	Message       string `json:"message"`
	Label         string `json:"label"`
}

// ClientPushLog client push data to room
//easyjson:json
type ClientPushLog struct {
	Topic   string `json:"topic"`
	MsgType uint8  `json:"msgType"`
	Message string `json:"message"`
	MsgID   uint16 `json:"msgID"`
}

// ClientPushAckLog client push ack log
//easyjson:json
type ClientPushAckLog struct {
	Topic         string `json:"topic"`
	MsgType       uint8  `json:"msgType"`
	OriginalMsgID string `json:"originalMsgID"`
	NewMsgID      uint16 `json:"newMsgID"`
}

// MessageLog message log
//easyjson:json
type MessageLog struct {
	Type       string `json:"type"`       //类型
	Topic      string `json:"resource"`   //主题
	Token      string `json:"token"`      //
	UserName   string `json:"opId"`       //用户名
	UserStatus string `json:"userStatus"` //用户状态
	MsgID      string `json:"msgId"`      //消息id
	Category   string `json:"category"`   //短号(消息ID截取前三位)
	Message    string `json:"content"`    //消息
}

//动作
const (
	MsgPush   string = "push"   //消息推送
	MsgAck    string = "ack"    //消息确认
	MsgExpire string = "expire" //消息过期
)
