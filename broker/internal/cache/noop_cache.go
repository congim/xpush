package cache

type NoopCache struct {
}

func (n *NoopCache) Init() error {
	return nil
}

func (n *NoopCache) Unsubscribe(userName string, topic string) error {
	return nil
}

func (n *NoopCache) Inc(string, string) error {
	return nil
}

func (n *NoopCache) StoreMsgID(string, string, string) error {
	return nil
}

func (n *NoopCache) Unread(string, string) (bool, error) {
	return false, nil
}

func newNoopCache() *NoopCache {
	return &NoopCache{}
}
