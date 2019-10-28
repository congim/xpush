package cache

import "log"

type NoopCache struct {
}

func (n *NoopCache) Init() error {
	return nil
}

func (n *NoopCache) Logout(string, string) error {
	return nil
}

func (n *NoopCache) Login(cid uint64, name string) error {
	log.Println(cid, name)
	return nil
}

func (n *NoopCache) GetBroker(uint64) (string, bool) {
	return "", false
}

func (n *NoopCache) Subscribe(userName string, topic string) error {
	return nil
}

func (n *NoopCache) PubCount(topic string, count int) error {
	return nil
}

func (n *NoopCache) Ack(userName string, topic string, count uint64) error {
	return nil
}

func (n *NoopCache) Unsubscribe(userName string, topic string) error {
	return nil
}

func newNoopCache() *NoopCache {
	return &NoopCache{}
}
