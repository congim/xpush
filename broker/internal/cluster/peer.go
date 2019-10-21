package cluster

import (
	"net/rpc"

	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type Peer struct {
	name   string
	addr   string
	logger *zap.Logger
	client *rpc.Client
}

func newPeer(addr, name string, logger *zap.Logger) (*Peer, error) {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		logger.Warn("rpc dial failed", zap.Error(err))
		return nil, err
	}

	p := &Peer{
		name:   name,
		addr:   addr,
		client: client,
		logger: logger,
	}

	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		p.OnMessage("world")
	//	}
	//}()

	return p, nil
}

func (p *Peer) Close() error {
	if p.client != nil {
		_ = p.client.Close()
	}
	return nil
}

func (p *Peer) OnMessage(msg *message.Message) (*message.Reply, error) {
	reply := message.NewReply()
	err := p.client.Call("RPCServer.OnMessage", msg, reply)
	if err != nil {
		p.logger.Warn("onmessage is failed", zap.Error(err))
		return reply, err
	}
	return reply, nil
}
