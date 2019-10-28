package cache

import (
	"github.com/congim/xpush/broker/internal/cache/redis"
	"github.com/congim/xpush/config"
	"go.uber.org/zap"
)

type Cache interface {
	Init() error
	GetBroker(uint64) (string, bool)
	Subscribe(string, string) error
	Unsubscribe(string, string) error
	PubCount(string, int) error
	Ack(string, string, uint64) error
	//Login(uint64, string) error
	//Logout(string) error
}

func New(conf *config.Cache, l *zap.Logger) Cache {
	if conf.Name == "redis" {
		return redis.New(conf.Redis, l)
	} else {
		return newNoopCache()
	}
}
