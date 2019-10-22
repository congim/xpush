package mem

import (
	"log"
	"sync"
	"sync/atomic"
)

type Mem struct {
	clients      sync.Map
	userTopics   sync.Map
	topicCounter sync.Map
	ackCounter   sync.Map
}

type receiveInfo struct {
	lastMsgID string
	received  uint64
}

func newreceiveInfo() *receiveInfo {
	return &receiveInfo{}
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

	return nil
}

func (m *Mem) Ack(userName string, topic string, msgID string) error {
	user, ok := m.ackCounter.Load(userName)
	if !ok {
		user = &sync.Map{}
		m.ackCounter.Store(userName, user)
	}
	topicInfo, ok := user.(*sync.Map).Load(topic)
	if !ok {
		topicInfo = newreceiveInfo()
		user.(*sync.Map).Store(topic, topicInfo)
	}

	topicInfo.(*receiveInfo).lastMsgID = msgID
	topicInfo.(*receiveInfo).received++

	log.Println(topicInfo.(*receiveInfo))
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

func New() *Mem {
	return &Mem{
		//clients:      &sync.Map{},
		//userTopics:   &sync.Map{},
		//topicCounter: &sync.Map{},
		//ackCounter:   &sync.Map{},
	}
}
