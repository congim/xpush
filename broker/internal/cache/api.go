package cache

import (
	"github.com/congim/xpush/broker/internal/cache/redis"
	"github.com/congim/xpush/config"
	"go.uber.org/zap"
)

// Cache ...
type Cache interface {
	Init() error
	Subscribe(string, string) error
	Unsubscribe(string, string) error
	Publish(string, string) error
	StoreMsgID(string, string, string) error
	// Ack(string, string, uint64) error
	// Login(string, string) error
	// PubCount(string, string, int) error
	// GetBroker(uint64) (string, bool)
	// UnRead(string, []string) (*message.UnRead, error)
	//Login(uint64, string) error
	//Logout(string) error
}

// New ...
func New(conf *config.Cache, l *zap.Logger) Cache {
	if conf.Name == "redis" {
		return redis.New(conf.Redis, l)
	} else {
		return newNoopCache()
	}
}
