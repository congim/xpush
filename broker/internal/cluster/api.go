package cluster

import (
	"time"

	"github.com/congim/xpush/config"
	"go.uber.org/zap"
)

type Cluster interface {
	Start() error
	Join() error
	Leave(time.Duration) error
	Close() error
}

func New(conf *config.Cluster, logger *zap.Logger, notify func(*Event) error) Cluster {
	swarm := swarm(conf, logger, notify)
	return swarm
}
