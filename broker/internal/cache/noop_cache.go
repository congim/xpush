package cache

import "log"

type NoopCache struct {
}

func (n *NoopCache) Init() error {
	return nil
}

func (n *NoopCache) Logout(cid uint64) error {
	log.Println(cid)
	return nil
}

func (n *NoopCache) Login(cid uint64, name string) error {
	log.Println(cid, name)
	return nil
}

func (n *NoopCache) Get(uint64) (string, bool) {
	return "", false
}

func newNoopCache() *NoopCache {
	return &NoopCache{}
}
