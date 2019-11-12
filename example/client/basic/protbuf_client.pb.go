// Code generated by protoc-gen-go.
// source: protbuf_client.proto
// DO NOT EDIT!

/*
Package basic is a generated protocol buffer package.

It is generated from these files:
	protbuf_client.proto

It has these top-level messages:
	ClientMsg
	MsgsPack
	Msgs
	CreateGroup
	GroupChatMsg
	FindGroup
	JoinGroup
	AgreeJoinGroup
	RefusedJoinGroup
	ExitGroup
	SetAdmin
	DelGroup
	KickGroup
	Response
	DeviceToken
	SetUser
	GroupUser
	GroupMsg
*/
package basic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientMsg struct {
	Msgid            *string `protobuf:"bytes,1,opt,name=msgid" json:"msgid,omitempty"`
	Mtype            *int32  `protobuf:"varint,3,req,name=mtype" json:"mtype,omitempty"`
	Time             *int64  `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	Nick             *string `protobuf:"bytes,4,opt,name=nick" json:"nick,omitempty"`
	Fid              *string `protobuf:"bytes,5,opt,name=fid" json:"fid,omitempty"`
	Ftopic           *string `protobuf:"bytes,6,opt,name=ftopic" json:"ftopic,omitempty"`
	Tid              *string `protobuf:"bytes,7,req,name=tid" json:"tid,omitempty"`
	Body             []byte  `protobuf:"bytes,8,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientMsg) Reset()                    { *m = ClientMsg{} }
func (m *ClientMsg) String() string            { return proto.CompactTextString(m) }
func (*ClientMsg) ProtoMessage()               {}
func (*ClientMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientMsg) GetMsgid() string {
	if m != nil && m.Msgid != nil {
		return *m.Msgid
	}
	return ""
}

func (m *ClientMsg) GetMtype() int32 {
	if m != nil && m.Mtype != nil {
		return *m.Mtype
	}
	return 0
}

func (m *ClientMsg) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *ClientMsg) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

func (m *ClientMsg) GetFid() string {
	if m != nil && m.Fid != nil {
		return *m.Fid
	}
	return ""
}

func (m *ClientMsg) GetFtopic() string {
	if m != nil && m.Ftopic != nil {
		return *m.Ftopic
	}
	return ""
}

func (m *ClientMsg) GetTid() string {
	if m != nil && m.Tid != nil {
		return *m.Tid
	}
	return ""
}

func (m *ClientMsg) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type MsgsPack struct {
	CompressMsg      *int32 `protobuf:"varint,1,req,name=compressMsg" json:"compressMsg,omitempty"`
	Body             []byte `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MsgsPack) Reset()                    { *m = MsgsPack{} }
func (m *MsgsPack) String() string            { return proto.CompactTextString(m) }
func (*MsgsPack) ProtoMessage()               {}
func (*MsgsPack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MsgsPack) GetCompressMsg() int32 {
	if m != nil && m.CompressMsg != nil {
		return *m.CompressMsg
	}
	return 0
}

func (m *MsgsPack) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// 数据打包列表
type Msgs struct {
	MsgList          []*ClientMsg `protobuf:"bytes,1,rep,name=MsgList" json:"MsgList,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Msgs) Reset()                    { *m = Msgs{} }
func (m *Msgs) String() string            { return proto.CompactTextString(m) }
func (*Msgs) ProtoMessage()               {}
func (*Msgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Msgs) GetMsgList() []*ClientMsg {
	if m != nil {
		return m.MsgList
	}
	return nil
}

// 创建群
type CreateGroup struct {
	Topic            *string `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Reqcheck         *int32  `protobuf:"varint,3,req,name=reqcheck" json:"reqcheck,omitempty"`
	Label            *string `protobuf:"bytes,4,opt,name=label" json:"label,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateGroup) Reset()                    { *m = CreateGroup{} }
func (m *CreateGroup) String() string            { return proto.CompactTextString(m) }
func (*CreateGroup) ProtoMessage()               {}
func (*CreateGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *CreateGroup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CreateGroup) GetReqcheck() int32 {
	if m != nil && m.Reqcheck != nil {
		return *m.Reqcheck
	}
	return 0
}

func (m *CreateGroup) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

//  群聊天信息
type GroupChatMsg struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Gnick            *string `protobuf:"bytes,2,req,name=gnick" json:"gnick,omitempty"`
	Body             *string `protobuf:"bytes,3,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GroupChatMsg) Reset()                    { *m = GroupChatMsg{} }
func (m *GroupChatMsg) String() string            { return proto.CompactTextString(m) }
func (*GroupChatMsg) ProtoMessage()               {}
func (*GroupChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GroupChatMsg) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *GroupChatMsg) GetGnick() string {
	if m != nil && m.Gnick != nil {
		return *m.Gnick
	}
	return ""
}

func (m *GroupChatMsg) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

// 群查找
type FindGroup struct {
	Gid              *string `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Describe         *string `protobuf:"bytes,2,opt,name=describe" json:"describe,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FindGroup) Reset()                    { *m = FindGroup{} }
func (m *FindGroup) String() string            { return proto.CompactTextString(m) }
func (*FindGroup) ProtoMessage()               {}
func (*FindGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FindGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *FindGroup) GetDescribe() string {
	if m != nil && m.Describe != nil {
		return *m.Describe
	}
	return ""
}

// 入群
type JoinGroup struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Postscript       *string `protobuf:"bytes,2,opt,name=postscript" json:"postscript,omitempty"`
	Nick             *string `protobuf:"bytes,3,opt,name=nick" json:"nick,omitempty"`
	Topic            *string `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	Opid             *string `protobuf:"bytes,5,opt,name=opid" json:"opid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JoinGroup) Reset()                    { *m = JoinGroup{} }
func (m *JoinGroup) String() string            { return proto.CompactTextString(m) }
func (*JoinGroup) ProtoMessage()               {}
func (*JoinGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *JoinGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *JoinGroup) GetPostscript() string {
	if m != nil && m.Postscript != nil {
		return *m.Postscript
	}
	return ""
}

func (m *JoinGroup) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

func (m *JoinGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *JoinGroup) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

// 同意入群
type AgreeJoinGroup struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Nick             *string `protobuf:"bytes,2,opt,name=nick" json:"nick,omitempty"`
	Topic            *string `protobuf:"bytes,3,req,name=topic" json:"topic,omitempty"`
	Opid             *string `protobuf:"bytes,4,req,name=opid" json:"opid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AgreeJoinGroup) Reset()                    { *m = AgreeJoinGroup{} }
func (m *AgreeJoinGroup) String() string            { return proto.CompactTextString(m) }
func (*AgreeJoinGroup) ProtoMessage()               {}
func (*AgreeJoinGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AgreeJoinGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *AgreeJoinGroup) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

func (m *AgreeJoinGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *AgreeJoinGroup) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

// 拒绝入群
type RefusedJoinGroup struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Postscript       *string `protobuf:"bytes,2,opt,name=postscript" json:"postscript,omitempty"`
	Nick             *string `protobuf:"bytes,3,opt,name=nick" json:"nick,omitempty"`
	Topic            *string `protobuf:"bytes,4,req,name=topic" json:"topic,omitempty"`
	Opid             *string `protobuf:"bytes,5,opt,name=opid" json:"opid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RefusedJoinGroup) Reset()                    { *m = RefusedJoinGroup{} }
func (m *RefusedJoinGroup) String() string            { return proto.CompactTextString(m) }
func (*RefusedJoinGroup) ProtoMessage()               {}
func (*RefusedJoinGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RefusedJoinGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *RefusedJoinGroup) GetPostscript() string {
	if m != nil && m.Postscript != nil {
		return *m.Postscript
	}
	return ""
}

func (m *RefusedJoinGroup) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

func (m *RefusedJoinGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *RefusedJoinGroup) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

// 退出群
type ExitGroup struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Opid             *string `protobuf:"bytes,2,opt,name=opid" json:"opid,omitempty"`
	Topic            *string `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExitGroup) Reset()                    { *m = ExitGroup{} }
func (m *ExitGroup) String() string            { return proto.CompactTextString(m) }
func (*ExitGroup) ProtoMessage()               {}
func (*ExitGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ExitGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *ExitGroup) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

func (m *ExitGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

// 设置管理员
type SetAdmin struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Opid             *string `protobuf:"bytes,2,req,name=opid" json:"opid,omitempty"`
	Topic            *string `protobuf:"bytes,3,req,name=topic" json:"topic,omitempty"`
	Action           *int32  `protobuf:"varint,4,req,name=action" json:"action,omitempty"`
	Nick             *string `protobuf:"bytes,5,opt,name=nick" json:"nick,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetAdmin) Reset()                    { *m = SetAdmin{} }
func (m *SetAdmin) String() string            { return proto.CompactTextString(m) }
func (*SetAdmin) ProtoMessage()               {}
func (*SetAdmin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetAdmin) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *SetAdmin) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

func (m *SetAdmin) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *SetAdmin) GetAction() int32 {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return 0
}

func (m *SetAdmin) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

// 删除群
type DelGroup struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Postscript       *string `protobuf:"bytes,2,opt,name=postscript" json:"postscript,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DelGroup) Reset()                    { *m = DelGroup{} }
func (m *DelGroup) String() string            { return proto.CompactTextString(m) }
func (*DelGroup) ProtoMessage()               {}
func (*DelGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DelGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *DelGroup) GetPostscript() string {
	if m != nil && m.Postscript != nil {
		return *m.Postscript
	}
	return ""
}

// 踢人
type KickGroup struct {
	Gid              *string `protobuf:"bytes,1,req,name=gid" json:"gid,omitempty"`
	Postscript       *string `protobuf:"bytes,2,opt,name=postscript" json:"postscript,omitempty"`
	Nick             *string `protobuf:"bytes,3,opt,name=nick" json:"nick,omitempty"`
	Topic            *string `protobuf:"bytes,4,req,name=topic" json:"topic,omitempty"`
	Opid             *string `protobuf:"bytes,5,req,name=opid" json:"opid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KickGroup) Reset()                    { *m = KickGroup{} }
func (m *KickGroup) String() string            { return proto.CompactTextString(m) }
func (*KickGroup) ProtoMessage()               {}
func (*KickGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *KickGroup) GetGid() string {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return ""
}

func (m *KickGroup) GetPostscript() string {
	if m != nil && m.Postscript != nil {
		return *m.Postscript
	}
	return ""
}

func (m *KickGroup) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

func (m *KickGroup) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *KickGroup) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

// 返回信息
type Response struct {
	Retcode          *int32  `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	ErrMsg           *string `protobuf:"bytes,2,opt,name=err_msg" json:"err_msg,omitempty"`
	Other            []byte  `protobuf:"bytes,3,opt,name=other" json:"other,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Response) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *Response) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *Response) GetOther() []byte {
	if m != nil {
		return m.Other
	}
	return nil
}

// 设备token
type DeviceToken struct {
	Token            *string `protobuf:"bytes,1,req,name=Token" json:"Token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeviceToken) Reset()                    { *m = DeviceToken{} }
func (m *DeviceToken) String() string            { return proto.CompactTextString(m) }
func (*DeviceToken) ProtoMessage()               {}
func (*DeviceToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DeviceToken) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

// 设置用户信息
type SetUser struct {
	Nick             *string `protobuf:"bytes,1,req,name=Nick" json:"Nick,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetUser) Reset()                    { *m = SetUser{} }
func (m *SetUser) String() string            { return proto.CompactTextString(m) }
func (*SetUser) ProtoMessage()               {}
func (*SetUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SetUser) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

// 群用户信息
type GroupUser struct {
	Sid              *string `protobuf:"bytes,1,req,name=sid" json:"sid,omitempty"`
	Opid             *string `protobuf:"bytes,2,req,name=opid" json:"opid,omitempty"`
	Nick             *string `protobuf:"bytes,3,opt,name=nick" json:"nick,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GroupUser) Reset()                    { *m = GroupUser{} }
func (m *GroupUser) String() string            { return proto.CompactTextString(m) }
func (*GroupUser) ProtoMessage()               {}
func (*GroupUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GroupUser) GetSid() string {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return ""
}

func (m *GroupUser) GetOpid() string {
	if m != nil && m.Opid != nil {
		return *m.Opid
	}
	return ""
}

func (m *GroupUser) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

// 群信息
type GroupMsg struct {
	Id               *string      `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Topic            *string      `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
	Nick             *string      `protobuf:"bytes,3,req,name=nick" json:"nick,omitempty"`
	Master           *GroupUser   `protobuf:"bytes,4,req,name=master" json:"master,omitempty"`
	Admins           []*GroupUser `protobuf:"bytes,5,rep,name=admins" json:"admins,omitempty"`
	Users            []*GroupUser `protobuf:"bytes,6,rep,name=users" json:"users,omitempty"`
	Other            []byte       `protobuf:"bytes,7,opt,name=other" json:"other,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GroupMsg) Reset()                    { *m = GroupMsg{} }
func (m *GroupMsg) String() string            { return proto.CompactTextString(m) }
func (*GroupMsg) ProtoMessage()               {}
func (*GroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GroupMsg) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *GroupMsg) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *GroupMsg) GetNick() string {
	if m != nil && m.Nick != nil {
		return *m.Nick
	}
	return ""
}

func (m *GroupMsg) GetMaster() *GroupUser {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *GroupMsg) GetAdmins() []*GroupUser {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *GroupMsg) GetUsers() []*GroupUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GroupMsg) GetOther() []byte {
	if m != nil {
		return m.Other
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientMsg)(nil), "basic.ClientMsg")
	proto.RegisterType((*MsgsPack)(nil), "basic.MsgsPack")
	proto.RegisterType((*Msgs)(nil), "basic.Msgs")
	proto.RegisterType((*CreateGroup)(nil), "basic.CreateGroup")
	proto.RegisterType((*GroupChatMsg)(nil), "basic.GroupChatMsg")
	proto.RegisterType((*FindGroup)(nil), "basic.FindGroup")
	proto.RegisterType((*JoinGroup)(nil), "basic.JoinGroup")
	proto.RegisterType((*AgreeJoinGroup)(nil), "basic.AgreeJoinGroup")
	proto.RegisterType((*RefusedJoinGroup)(nil), "basic.RefusedJoinGroup")
	proto.RegisterType((*ExitGroup)(nil), "basic.ExitGroup")
	proto.RegisterType((*SetAdmin)(nil), "basic.SetAdmin")
	proto.RegisterType((*DelGroup)(nil), "basic.DelGroup")
	proto.RegisterType((*KickGroup)(nil), "basic.KickGroup")
	proto.RegisterType((*Response)(nil), "basic.Response")
	proto.RegisterType((*DeviceToken)(nil), "basic.DeviceToken")
	proto.RegisterType((*SetUser)(nil), "basic.SetUser")
	proto.RegisterType((*GroupUser)(nil), "basic.GroupUser")
	proto.RegisterType((*GroupMsg)(nil), "basic.GroupMsg")
}

var fileDescriptor0 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xdf, 0x8f, 0xd2, 0x40,
	0x10, 0x0e, 0x2d, 0x85, 0x76, 0xc0, 0x93, 0x54, 0x13, 0xfb, 0x60, 0x22, 0xf6, 0x09, 0x35, 0xf2,
	0x60, 0x34, 0x31, 0xf1, 0xe9, 0xc2, 0xf9, 0x23, 0x9e, 0x1a, 0x83, 0x1a, 0xdf, 0xbc, 0x94, 0xed,
	0x00, 0x1b, 0xa0, 0x5b, 0x77, 0x17, 0xe3, 0x3d, 0xf8, 0xf7, 0xf8, 0x6f, 0x3a, 0xbb, 0x4b, 0x8b,
	0x1c, 0x90, 0xdc, 0xc3, 0x3d, 0xdd, 0xcd, 0xce, 0x37, 0x33, 0xdf, 0xf7, 0x4d, 0x07, 0xb8, 0x5b,
	0x4a, 0xa1, 0x27, 0xeb, 0xe9, 0x05, 0x5b, 0x72, 0x2c, 0xf4, 0xd0, 0x84, 0x22, 0x0e, 0x26, 0x99,
	0xe2, 0x2c, 0xfd, 0x03, 0xd1, 0xc8, 0x3e, 0x7f, 0x54, 0xb3, 0xf8, 0x16, 0x04, 0x2b, 0x35, 0xe3,
	0x79, 0xd2, 0xe8, 0x37, 0x06, 0x91, 0x0d, 0xf5, 0x65, 0x89, 0x89, 0xdf, 0xf7, 0x06, 0x41, 0xdc,
	0x85, 0xa6, 0xe6, 0x2b, 0x4c, 0x3c, 0x8a, 0x7c, 0x13, 0x15, 0x9c, 0x2d, 0x92, 0xa6, 0x85, 0x76,
	0xc0, 0x9f, 0x52, 0x5d, 0x60, 0x83, 0x13, 0x68, 0x4d, 0xb5, 0x28, 0x39, 0x4b, 0x5a, 0x55, 0x52,
	0x53, 0xb2, 0x4d, 0x75, 0x91, 0xa9, 0x9b, 0x88, 0xfc, 0x32, 0x09, 0x29, 0xea, 0xa6, 0x4f, 0x21,
	0xa4, 0xc1, 0xea, 0x73, 0xc6, 0x16, 0xf1, 0x1d, 0xe8, 0x30, 0xb1, 0x2a, 0x25, 0x2a, 0x45, 0x6f,
	0xc4, 0x61, 0x33, 0xd4, 0xc2, 0x3d, 0x0b, 0x7f, 0x04, 0x4d, 0x03, 0x8f, 0x1f, 0x42, 0x9b, 0xfe,
	0x7e, 0xe0, 0x4a, 0x13, 0xcc, 0x1f, 0x74, 0x9e, 0xf5, 0x86, 0x56, 0xce, 0xb0, 0xd6, 0x92, 0x9e,
	0x43, 0x67, 0x24, 0x31, 0xd3, 0xf8, 0x56, 0x8a, 0x75, 0x69, 0xb4, 0x38, 0x4a, 0x8d, 0x8a, 0x45,
	0x91, 0x6d, 0xb4, 0x44, 0x71, 0x0f, 0x42, 0x89, 0x3f, 0xd9, 0x1c, 0x49, 0x8f, 0xd3, 0x4a, 0xf0,
	0x65, 0x36, 0xc1, 0xa5, 0x93, 0x97, 0xbe, 0x84, 0xae, 0x6d, 0x33, 0x9a, 0x67, 0xd6, 0x28, 0x52,
	0xe4, 0x6c, 0xf2, 0x9c, 0x4d, 0x33, 0x6b, 0x85, 0xb7, 0x23, 0xd0, 0x34, 0x8a, 0xd2, 0xc7, 0x10,
	0xbd, 0xe1, 0x45, 0xee, 0x48, 0xd4, 0x65, 0x0d, 0x37, 0x34, 0x47, 0xc5, 0x24, 0x9f, 0x18, 0x1a,
	0x66, 0xca, 0x77, 0x88, 0xde, 0x0b, 0x5e, 0x5c, 0xc1, 0x9a, 0x9e, 0x31, 0x40, 0x29, 0x94, 0x36,
	0xe8, 0x52, 0x3b, 0x74, 0xbd, 0x00, 0xbf, 0xda, 0x95, 0xd3, 0xd7, 0xac, 0x92, 0x14, 0x6d, 0x16,
	0x92, 0xbe, 0x83, 0x93, 0xd3, 0x99, 0x44, 0x3c, 0xd2, 0xbd, 0xea, 0xe4, 0xed, 0x76, 0xf2, 0xab,
	0xa4, 0xed, 0xd4, 0xb4, 0x72, 0x7e, 0x40, 0x6f, 0x8c, 0xd3, 0xb5, 0xc2, 0xfc, 0x26, 0x98, 0x7a,
	0x7b, 0x4c, 0x5f, 0x40, 0xf4, 0xfa, 0x37, 0xd7, 0x87, 0x49, 0x5a, 0xdc, 0x1e, 0x49, 0x53, 0x36,
	0x86, 0xf0, 0x0b, 0xea, 0xd3, 0x7c, 0xc5, 0x8b, 0x63, 0x55, 0xde, 0xbe, 0x34, 0xfa, 0x4e, 0x33,
	0xa6, 0xb9, 0x28, 0x2c, 0x95, 0xa0, 0xe6, 0xe9, 0xa8, 0x3c, 0x81, 0xf0, 0x0c, 0x97, 0xd7, 0x93,
	0x68, 0x56, 0x77, 0x4e, 0xa5, 0x37, 0x6b, 0x88, 0x31, 0xfc, 0x15, 0x84, 0x63, 0x54, 0xa5, 0x28,
	0x14, 0xc6, 0xb7, 0xa1, 0x2d, 0x51, 0x33, 0x91, 0xe3, 0xe6, 0x38, 0xe8, 0x01, 0xa5, 0xbc, 0xa0,
	0x9b, 0xdd, 0xda, 0x22, 0xf4, 0x1c, 0xa5, 0xed, 0xdc, 0x4d, 0xef, 0x43, 0xe7, 0x0c, 0x7f, 0x71,
	0x86, 0x5f, 0xc5, 0x02, 0x0b, 0x93, 0xb5, 0xff, 0x38, 0x66, 0xe9, 0x3d, 0x68, 0x93, 0x69, 0xdf,
	0x14, 0x4a, 0x33, 0xf3, 0x93, 0x21, 0xe4, 0x12, 0xcf, 0x21, 0xb2, 0x42, 0x6c, 0x8a, 0xc4, 0xa8,
	0x23, 0x76, 0xee, 0xc8, 0x48, 0xff, 0x36, 0x20, 0xb4, 0x65, 0xe6, 0x40, 0x00, 0xbc, 0xff, 0xef,
	0xc3, 0xe9, 0xbb, 0x5a, 0x65, 0xa2, 0x3e, 0xb4, 0x56, 0x99, 0xd2, 0x44, 0xd9, 0xa8, 0xdf, 0x1e,
	0xf2, 0x96, 0x00, 0x21, 0x32, 0xb3, 0x58, 0x45, 0x8e, 0xf8, 0x07, 0x11, 0x0f, 0x20, 0xa0, 0x2f,
	0x52, 0x2a, 0xfa, 0xb9, 0x39, 0x0c, 0xa8, 0x6d, 0x69, 0x1b, 0x5b, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x48, 0xa4, 0xdb, 0x11, 0x05, 0x00, 0x00,
}