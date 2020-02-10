package cluster

import (
	"net"
	"net/rpc"

	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type RPCServer struct {
	name   string
	notify func(*message.Event) error
}

func (r *RPCServer) SyncMessage(msg *message.Message, reply *message.Reply) error {
	event := &message.Event{
		Msgs: []*message.Message{msg},
		Type: msg.Type,
	}

	return r.notify(event)
}

func startRPC(addr string, logger *zap.Logger, notify func(*message.Event) error) error {
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
