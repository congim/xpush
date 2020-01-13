package cluster

import (
	"net"
	"net/rpc"

	"go.uber.org/zap"
)

type RPCServer struct {
	name   string
	notify func(*Event) error
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
