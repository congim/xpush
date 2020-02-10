package cache

type NoopCache struct {
}

func (n *NoopCache) Init() error {
	return nil
}

func (n *NoopCache) Unsubscribe(userName string, topic string) error {
	return nil
}

func (n *NoopCache) Incr(string) error {
	return nil
}

func (n *NoopCache) StoreMsgID(string, string, string) error {
	return nil
}

func (n *NoopCache) GetIncr(string) (int, error) {
	return 0, nil
}

func newNoopCache() *NoopCache {
	return &NoopCache{}
}
