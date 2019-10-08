package cluster

import (
	"github.com/hashicorp/memberlist"
	"go.uber.org/zap"
)

type EventType int

const (
	_ EventType = iota
	Join
	Leave
	Update
)

type Event struct {
	Type EventType
	Name string
	Addr string
	Port uint16
}

type event struct {
	c      *cluster
	logger *zap.Logger
}

func newEvent(c *cluster, logger *zap.Logger) *event {
	return &event{
		c:      c,
		logger: logger,
	}
}

func (e *event) NotifyJoin(n *memberlist.Node) {
	event := &Event{
		Type: Join,
		Name: n.Name,
		Addr: n.Addr.String(),
		Port: n.Port,
	}
	if err := e.c.notify(event); err != nil {
		e.logger.Warn("notify failed", zap.String("type", "Join"), zap.Any("event", event))
	}
}

func (e *event) NotifyLeave(n *memberlist.Node) {
	event := &Event{
		Type: Leave,
		Name: n.Name,
		Addr: n.Addr.String(),
		Port: n.Port,
	}
	if err := e.c.notify(event); err != nil {
		e.logger.Warn("notify failed", zap.String("type", "Leave"), zap.Any("event", event))
	}
}

func (e *event) NotifyUpdate(n *memberlist.Node) {
	event := &Event{
		Type: Update,
		Name: n.Name,
		Addr: n.Addr.String(),
		Port: n.Port,
	}
	if err := e.c.notify(event); err != nil {
		e.logger.Warn("notify failed", zap.String("type", "Update"), zap.Any("event", event))
	}
}
