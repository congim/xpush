package cluster

import (
	"context"

	"github.com/congim/xpush/broker/internal/proto/peer"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Peer struct {
	name   string
	addr   string
	logger *zap.Logger
	client peer.PeerClient
	conn   *grpc.ClientConn
}

func newPeer(addr, name string, logger *zap.Logger) (*Peer, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.Error("grp dial failed", zap.Error(err))
		return nil, err
	}

	p := &Peer{
		name:   name,
		addr:   addr,
		logger: logger,
		conn:   conn,
	}
	p.client = peer.NewPeerClient(conn)
	return p, nil
}

func (p *Peer) Close() error {
	if p.conn != nil {
		_ = p.conn.Close()
	}
	return nil
}

func (p *Peer) OnMessage(msg *peer.Message) (*peer.Reply, error) {
	reply, err := p.client.OnMessage(context.Background(), msg)
	if err != nil {
		p.logger.Warn("onmessage is failed", zap.Error(err))
		return reply, err
	}
	return reply, nil
}
