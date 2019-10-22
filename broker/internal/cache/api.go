package cache

import (
	"github.com/congim/xpush/broker/internal/cache/mem"
	"github.com/congim/xpush/config"
	"go.uber.org/zap"
)

type Cache interface {
	Init() error
	Login(uint64, string) error
	Logout(uint64) error
	Get(uint64) (string, bool)
	Subscribe(string, string) error
	PubCount(string, int) error
}

func New(conf *config.Cache, l *zap.Logger) Cache {
	if conf.Name == "mem" {
		return mem.New()
	} else {
		return newNoopCache()
	}
}
