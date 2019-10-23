package mem

import (
	"sync"
	"sync/atomic"
)

type Mem struct {
	clients      sync.Map
	userTopics   sync.Map
	topicCounter sync.Map
	ackCounter   sync.Map
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

func (m *Mem) Subscribe(userName string, topic string) error {
	user, ok := m.userTopics.Load(userName)
	if !ok {
		user = &sync.Map{}
		m.userTopics.Store(userName, user)
	}

	user.(*sync.Map).Store(topic, 0)

	// @TODO 订阅时要将当前用户已接收和最近msgID和topic当前msgid、发送量保持一致

	return nil
}

func (m *Mem) Ack(userName string, topic string, count uint64) error {
	user, ok := m.ackCounter.Load(userName)
	if !ok {
		user = &sync.Map{}
		m.ackCounter.Store(userName, user)
	}
	counter, ok := user.(*sync.Map).Load(topic)
	if !ok {
		var tmpCount uint64
		counter = &tmpCount
		user.(*sync.Map).Store(topic, counter)
	}
	atomic.AddUint64(counter.(*uint64), count)
	return nil
}

func (m *Mem) PubCount(topic string, count int) error {
	counter, ok := m.topicCounter.Load(topic)
	if !ok {
		var tmpCount uint64
		counter = &tmpCount

		m.topicCounter.Store(topic, counter)
	}
	atomic.AddUint64(counter.(*uint64), uint64(count))
	return nil
}

func (m *Mem) Unsubscribe(userName string, topic string) error {
	user, ok := m.userTopics.Load(userName)
	if ok {
		user.(*sync.Map).Delete(topic)
	}

	// @TODO 客户端已接收信息保存清空
	ackUser, ok := m.ackCounter.Load(userName)
	if !ok {
		return nil
	}
	ackUser.(*sync.Map).Delete(topic)
	return nil
}

func New() *Mem {
	return &Mem{
		//clients:      &sync.Map{},
		//userTopics:   &sync.Map{},
		//topicCounter: &sync.Map{},
		//ackCounter:   &sync.Map{},
	}
}
