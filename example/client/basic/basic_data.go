package basic

import "fmt"

//PushArgs cache center Communication data
type PushArgs struct {
	TopicList []string
	UserList  []string
	MsgID     string
	MsgType   int
	Message   []byte //package、序列化之后的数据
	Label     string
}

// SessionMsg session
type SessionMsg struct {
	Sid    string
	Groups []string
}

// SubTopicMsg subscribe topic
type SubTopicMsg struct {
	Rip        string
	NatsTopic  string
	Opid       string
	TopicList  []TopicMsg
	MayorTopic string
	SystemType int
}

// Show show msg
func (sm *SubTopicMsg) Show() {
	fmt.Println("[DEBUG] [Show] - Rip", sm.Rip)
	fmt.Println("[DEBUG] [Show] - Opid", sm.Opid)
	for k, v := range sm.TopicList {
		fmt.Println("[DEBUG] [Show] - TopicList", k, v)
	}
	fmt.Println("[DEBUG] [Show] - MayorTopic", sm.MayorTopic)
}

// InitSession 初始用户
type InitSession struct {
	Rip                string
	Opid               string
	MayorTopic         string
	DeviceLabel        string // 设备标签为主topic后四为
	OriginalMayorTopic string // 原始主topic
	Token              string // 令牌自生成
	ConnType           int    // 客户端连接方式
	ClientId           string //客户端连接 clientid
	SystemType         int
}

// SetToken 自生成token令牌
func (is *InitSession) SetToken() {
	is.Token = Token(is.MayorTopic)
}

// DelTopicMsg DelTopicMsg
type DelTopicMsg struct {
	Sid       string
	TopicList []TopicMsg
}

// AckMsg 消息确认
type AckMsg struct {
	Topic   string
	MsgID   string
	MsgType int
	Sid     string
	MsgIDs  []string
}

// Show show msg
func (am *AckMsg) Show() {
	fmt.Println("[DEBUG] Topic:", string(am.Topic))
	fmt.Println("[DEBUG] MsgID:", string(am.MsgID))
	fmt.Println("[DEBUG] MsgType:", am.MsgType)
	fmt.Println("Sid:", string(am.Sid))
	for _, v := range am.MsgIDs {
		fmt.Println("[DEBUG] MsgIDs:", v)
	}
}

// CleanSession room清除cache客户信息结构
type CleanSession struct {
	MayorTopic string
	Rip        string
}

// LogOutMessage 踢出链接
type LogOutMessage struct {
	SessionID string
	Token     string
	ClientID  string
}

// TopicMsg topic信息 在订阅主题的时候发送的信息，可以区主题类别，以便处理
type TopicMsg struct {
	MsgType int
	Topic   string
}

// MsgIDInfo 用来映射客户端topic和业方的MsgID
type MsgIDInfo struct {
	Topic  string
	MsgIDs []string
}

//动作
const (
	MainTopic        string = "-m" //主topic
	UnicastTopic     string = "-s" //单推送
	GroupleTopic     string = "-g" //群推送
	BroadCastTopic   string = "-b" //公共广播推送
	PrivateChatTopic string = "-p" //私聊
)

// 主topic
const (
	ActionMainPush int = 1000 // 主
)

// 私聊
const (
	ActionPrivateChatPush         int = 2000 // 私聊
	ActionPrivateChatFileTransfer int = 2001 // 文件协议
)

// 单播
const (
	ActionUnicastPush         int = 3000 // 单播
	ActionUnicastFileTransfer int = 3001 // 文件协议
)

// 群聊
const (
	ActionGrouplePush      int = 4000 // 群聊
	ActionCreateGroup      int = 4001 // 创建群申请
	ActionJoinGroup        int = 4002 // 入群申请
	ActionAgreeJoinGroup   int = 4003 // 同意入群申请
	ActionRefusedJoinGroup int = 4004 // 拒绝入群申请
	ActionExitGroup        int = 4005 // 退群
	ActionDeleteGroup      int = 4006 // 删除群申请
	ActionKickGroup        int = 4007 // 踢出群
	ActionFindGroup        int = 4008 // 查找群
	ActionGetGruopMsg      int = 4009 // 获取群信息
	NoticSubGroup          int = 4010 // 推送给客户端订阅组消息
	ActionSetAdmin         int = 4011 // 设置管理员
	// ActionCreateGroupSucess int = 4011 // 创建群成功

)

// 公共广播
const (
	ActionBroadCastPush         int = 5000 // 广播
	ActionBroadCastFileTransfer int = 5001 // 文件协议
)

// 用户设置
const (
	ActionSetUser     int = 6000 // set user
	ActionDeviceToken int = 6001 // iostoken
	ActionRepeatLogin int = 6002 // 账号被顶
)

// 上报信息
const (
	ActionTotalMsg int = 6500 // 接收到信息
	ActionAckMsg   int = 6501 // 消费信息
	ActionExpire   int = 6502 // 过期消息

	ActionLogIn  int = 6503 // 登陆
	ActionLogOut int = 6504 // 登出
)

// 平台信息
const (
	PlatformXmpp string = "Xmpp" // Xmpp 平台
	PlatformMqtt string = "Mqtt" // Mqtt 平台
)

//动作
const (
	// ActionLogOut     uint8 = 6 // 登出
	ActionSelect     uint8 = 5 // 查询
	ActionPushMsg    uint8 = 7 // 统一推送给客户端
	ResponseERR      int32 = 1
	ResponseOK       int32 = 0
	ActionGzip       int32 = 1 //Gzip
	ActionSnappy     int32 = 2 //Snappy
	ActionCompress   int32 = 1 //压缩
	ActionUnCompress int32 = 0 //未压缩
)

const (
//ConnTypeTCP   int    = 8000
//ConnTypeTLS   int    = 8001
//ConnTypeWEB   int    = 8002
//Certificate   int    = 8003
//TCP_Subscribe string = "Conn_TCP"
//TLS_Subscribe string = "Conn_TLS"
//WEB_Subscribe string = "Conn_WEB"
)

// 加群验证
const (
	//JoinCheckAnyone         int32  = 1              // 运行任何人
	//JoinCheckAuthentication int32  = 2              // 需身份验证
	//JoinCheckNever          int32  = 3              // 不允许任何人
	ApnsTokenKey string = "ApnsTokenKey" // apns token key
	//AddAdmin                int32  = 1              // 设置管理员
	//DelAdmin                int32  = 2              // 取消管理员
)

//设备类型
const (
//SYSTEM_ANDROID = 1 //android
//SYSTEM_IOS     = 2 //ios
//SYSTEM_WEB     = 3 //web
//SYSTEM_PC      = 4 //pc-windows
//SYSTEM_WEIXIn  = 5 //weixin公众号
)
