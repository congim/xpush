package cluster

import (
	"time"

	cluster "github.com/congim/xpush/broker/internal/cluster/swarm"
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type Cluster interface {
	Start() error
	Join() error
	Leave(time.Duration) error
	Close() error
	SyncMsg(*message.Message) ([]*message.Reply, error)
}

func New(conf *config.Cluster, logger *zap.Logger, notify func(*message.Event) error) Cluster {
	return cluster.New(conf, logger, notify)
}
