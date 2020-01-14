package cluster

import (
	"time"

	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type Cluster interface {
	Start() error
	Join() error
	Leave(time.Duration) error
	Close() error
	SyncMessage(*message.Message) ([]*message.Reply, error)
}

func New(conf *config.Cluster, logger *zap.Logger, notify func(*Event) error) Cluster {
	swarm := newSwarm(conf, logger, notify)
	return swarm
}
