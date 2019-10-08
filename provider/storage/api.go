package storage

import "github.com/congim/xpush/provider/storage/cassandra"

type Storage interface {
	Store() error
}

func New(storageName string) Storage {
	if storageName == "cassandra" {
		return cassandra.New()
	}
	return newNoopStorage()
}
