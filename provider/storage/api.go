package storage

import (
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"github.com/congim/xpush/provider/storage/mysql"
	"go.uber.org/zap"
)

type Storage interface {
	Init() error
	Store([]*message.Message) error
	Close() error
	Get(string, []byte, int) ([]*message.Message, error)
}

func New(conf *config.Storage, logger *zap.Logger) Storage {
	if conf.Name == "cassandra" {
		//
	} else if conf.Name == "fdb" {
		//return foundationdb.New(conf.Fdb, logger)
	} else if conf.Name == "mysql" {
		return mysql.New(conf.Mysql, logger)
	}
	return newNoopStorage(conf, logger)
}
