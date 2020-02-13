package message

type MsgControl byte

const (
	//_             byte = iota
	ClusterJoin   byte = 1 // 加入集群
	ClusterLeave       = 2 // 离开集群
	ClusterUpdate      = 3 // 更新(集群未启用)
	MsgPub             = 4 // 消息推送
	MsgPull            = 5 // 拉取消息
	Sub                = 6 // 订阅
	UnSub              = 7 // 取消订阅
	MsgUnread          = 8 // 有未读消息
)

//type MsgVersion byte
//
const (
	NoCompress byte = 0
	Compress        = 1
)

const (
	MAX_MESSAGE_PULL_COUNT int = 200
)

const (
	//Topic_Msg_Count string = "_topic_msg_count"
	//User_Msg_Count  string = "_user_msg_count"

	Topic_Msg_InsertTime string = "_topic_msg_insert_time"
	User_Msg_AckTime     string = "_user_msg_ack_time"
)

const (
	HASE_UNREAD_MSG bool = true
	NO_UNREAD_MSG        = false
)

var (
	MSG_NEWEST_OFFSET = []byte("0")
)
