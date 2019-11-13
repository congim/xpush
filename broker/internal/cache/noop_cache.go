package cache

type NoopCache struct {
}

func (n *NoopCache) Init() error {
	return nil
}

// func (n *NoopCache) Logout(string, string) error {
// 	return nil
// }

// func (n *NoopCache) Login(cid uint64, name string) error {
// 	log.Println(cid, name)
// 	return nil
// }

// func (n *NoopCache) GetBroker(uint64) (string, bool) {
// 	return "", false
// }

// func (n *NoopCache) Subscribe(userName string, topic string) error {
// 	return nil
// }

// func (n *NoopCache) PubCount(string, topic string, count int) error {
// 	return nil
// }

// func (n *NoopCache) Ack(userName string, topic string, count uint64) error {
// 	return nil
// }

func (n *NoopCache) Unsubscribe(userName string, topic string) error {
	return nil
}

// func (n *NoopCache) UnRead(userName string, topics []string) (*message.Unread, error) {
// 	return nil, nil
// }

// func (n *NoopCache) Login(userName string, brokerName string) error {
// 	return nil
// }

func (n *NoopCache) Publish(string, string) error {
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
