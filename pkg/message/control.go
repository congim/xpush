package message

type MsgControl byte

const (
	_             byte = iota
	ClusterJoin        // 加入集群
	ClusterLeave       // 离开集群
	ClusterUpdate      // 更新(集群未启用)
	MsgPub             // 消息推送
	MsgPull            // 拉取消息
	Sub                // 订阅
	UnSub              // 取消订阅
	MsgUnread          // 有未读消息
)

//type MsgVersion byte
//
const (
	_ byte = iota
	NoCompress
	Compress
)

const (
	MAX_MESSAGE_PULL_COUNT int = 200
)

const (
	Topic_Msg_Count string = "_topic_msg_count"
	User_Msg_Count  string = "_user_msg_count"
)

var (
	MSG_NEWEST_OFFSET = []byte("0")
)
