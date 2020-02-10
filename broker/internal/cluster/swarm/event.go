package cluster

import (
	"github.com/congim/xpush/pkg/message"
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type event struct {
	s      *Swarm
	logger *zap.Logger
}

func newEvent(s *Swarm, logger *zap.Logger) *event {
	return &event{
		s:      s,
		logger: logger,
	}
}

func (e *event) NotifyJoin(n *memberlist.Node) {
	event := &message.Event{
		Type: message.ClusterJoin,
		Name: n.Name,
		Addr: n.Addr.String(),
		Port: n.Port,
	}

	if n.Name != e.s.conf.Name {
		p, err := newPeer(n.Name, n.Name, e.logger)
		if err != nil {
			e.logger.Warn("new peer failed", zap.String("type", "Join"), zap.Any("event", event))
			return
		}

		if oldPeer, ok := e.s.peers.Load(n.Name); ok {
			_ = oldPeer.(*Peer).Close()
			e.s.peers.Delete(n.Name)
		}
		e.s.peers.Store(n.Name, p)
	}

	if err := e.s.notify(event); err != nil {
		e.logger.Warn("notify failed", zap.String("type", "Join"), zap.Any("event", event))
	}
}

func (e *event) NotifyLeave(n *memberlist.Node) {
	event := &message.Event{
		Type: message.ClusterLeave,
		Name: n.Name,
		Addr: n.Addr.String(),
		Port: n.Port,
	}

	if err := e.s.notify(event); err != nil {
		e.logger.Warn("notify failed", zap.String("type", "Leave"), zap.Any("event", event))
	}

	// 删除peer
	if peer, ok := e.s.peers.Load(n.Name); ok {
		_ = peer.(*Peer).Close()
		e.s.peers.Delete(n.Name)
	}
}

func (e *event) NotifyUpdate(n *memberlist.Node) {
	//event := &Event{
	//	Type: Update,
	//	Name: n.Name,
	//	Addr: n.Addr.String(),
	//	Port: n.Port,
	//}
	//if err := e.s.notify(event); err != nil {
	//	e.logger.Warn("notify failed", zap.String("type", "Update"), zap.Any("event", event))
	//}
}
