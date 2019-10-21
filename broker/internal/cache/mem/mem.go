package mem

import (
	"sync"
)

type Mem struct {
	clients sync.Map
}

func (m *Mem) Init() error {
	return nil
}

func (m *Mem) Logout(cid uint64) error {
	m.clients.Delete(cid)
	return nil
}

func (m *Mem) Login(cid uint64, clusterName string) error {
	m.clients.Store(cid, clusterName)
	return nil
}

func (m *Mem) Get(cid uint64) (string, bool) {
	value, ok := m.clients.Load(cid)
	if ok {
		return value.(string), true
	}
	return "", false
}

func New() *Mem {
	return &Mem{}
}
