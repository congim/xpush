package cluster

import (
	"net"
	"net/rpc"

	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type RPCServer struct {
	name   string
	notify func(*Event) error
}

func (r *RPCServer) OnMessage(msg *message.Message, reply *message.Reply) error {
	event := &Event{
		Msgs: []*message.Message{msg},
	}

	// 外部type和内部type转换一下
	if msg.Type == message.MsgPub {
		event.Type = Pub
	} else if msg.Type == message.Sub {
		event.Type = Sub
		return nil
	} else if msg.Type == message.UnSub {
		event.Type = UnSub
		return nil
	}

	return r.notify(event)
}

func startRPC(addr string, logger *zap.Logger, notify func(*Event) error) error {
	s := &RPCServer{
		name:   addr,
		notify: notify,
	}

	if err := rpc.Register(s); err != nil {
		logger.Error("rpc register failed", zap.Error(err))
		return err
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		logger.Error("net resolve tcp addr failed", zap.Error(err))
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
