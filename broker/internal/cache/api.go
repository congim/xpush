package cache

import (
	"github.com/congim/xpush/broker/internal/cache/redis"
	"github.com/congim/xpush/config"
	"go.uber.org/zap"
)

// Cache ...
type Cache interface {
	Init() error
	Unsubscribe(string, string) error
	StoreMsgID(string, string, string) error
	Publish(string, string) error
	Unread(string, string) (bool, error)
}

// New ...
func New(conf *config.Cache, l *zap.Logger) Cache {
	if conf.Name == "redis" {
		return redis.New(conf.Redis, l)
	}
	return newNoopCache()
}
