package storage

import (
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type noopStorage struct {
}

var _ Storage = (*noopStorage)(nil)

func newNoopStorage(conf *config.Storage, logger *zap.Logger) *noopStorage {
	return &noopStorage{}
}

func (n *noopStorage) Init() error {
	return nil
}

func (n *noopStorage) Store([]*message.Message, []string) error {
	return nil
}

func (n *noopStorage) Close() error {
	return nil
}

func (n *noopStorage) Get(string, int, int64) ([]*message.Message, error) {
	return nil, nil
}
